package service

import (
	"context"
	"crypto/sha512"

	"code.gopub.tech/errors"
	"code.gopub.tech/pub/dal/caches"
	"code.gopub.tech/pub/dal/model"
	"code.gopub.tech/pub/dal/query"
	"code.gopub.tech/pub/dto"
	"code.gopub.tech/pub/util"
	"github.com/youthlin/t"
	"gorm.io/gorm"
)

var cache = caches.GetDefaultCache()

// Installed 是否已经安装
func Installed(ctx context.Context) bool {
	val, err := cache.GetOrFetch(ctx, caches.Installed, func(ctx context.Context, key string) (any, error) {
		return queryInstalled(ctx), nil
	})
	if err != nil {
		return false
	}
	return val.(bool)
}

func queryInstalled(ctx context.Context) bool {
	o := query.Option
	do := o.WithContext(ctx)
	option, err := do.Where(o.Name.Eq(model.OptionNameInstalled)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return option.Value == model.OptionValueYes
}

// 更新安装状态缓存
func SetInstalled(ctx context.Context) {
	cache.Set(ctx, caches.Installed, true)
}

// GetStaticSalt 获取站点静态 salt
// 优先从缓存取，没有则查库，db 里没有就创建
func GetStaticSalt(ctx context.Context) (string, error) {
	val, err := cache.GetOrFetch(ctx, caches.StaticSalt, func(ctx context.Context, key string) (any, error) {
		o := query.Option
		option, err := o.WithContext(ctx).
			Attrs(o.Value.Value(util.RandStr(16))).
			Where(o.Name.Eq(model.OptionNameSalt)).
			FirstOrCreate()
		if err != nil {
			return "", errors.Wrapf(err, "failed to query/create option %v", model.OptionNameSalt)
		}
		return option.Value, nil
	})
	if err != nil {
		return "", errors.Wrapf(err, "failed to get salt")
	}
	return val.(string), nil
}

// Install 执行安装动作
func Install(ctx context.Context, req *dto.InstallReq) error {
	salt, err := GetStaticSalt(ctx)
	if err != nil {
		return err
	}
	if req.Salt != salt {
		return errors.Errorf("invalid request, mismatch salt")
	}
	t := t.WithContext(ctx)
	return query.Q.Transaction(func(tx *query.Query) error {
		u := tx.User
		salt := util.RandStr(16) // 每个用户在后端再单独随机加盐
		pass := sha512.Sum512([]byte(req.Password + salt))
		user := &model.User{
			Username: req.Username,
			Display:  req.Username,
			Email:    req.Email,
			Password: pass[:],
			Salt:     salt,
		}
		if err := u.WithContext(ctx).Create(user); err != nil {
			return errors.Wrapf(err, "failed to create user")
		}
		m := tx.UserMeta
		if err := m.WithContext(ctx).Create(&model.UserMeta{
			UserID: user.ID,
			Key:    model.UserMetaKeyRole,
			Value:  model.RoleSuperAdmin,
		}); err != nil {
			return errors.Wrapf(err, "failed to create user meta")
		}
		o := tx.Option
		option := &model.Option{Name: model.OptionNameSiteTitle, Value: req.SiteTitle}
		if err := o.WithContext(ctx).Save(option); err != nil {
			return errors.Wrapf(err, "failed to save site title")
		}
		option = &model.Option{Name: model.OptionNameInstalled, Value: model.OptionValueYes}
		if err := o.WithContext(ctx).Save(option); err != nil {
			return errors.Wrapf(err, "failed to update system option")
		}
		post := &model.Post{
			AuthorID: user.ID,
			Title:    t.T("Hello, World"),
			Content:  t.T(`This is a sample post. Go to <a href="/admin">admin page</a> to manage your site.`),
			Status:   model.PostStatusPublish,
		}
		if err := tx.Post.WithContext(ctx).Create(post); err != nil {
			return errors.Wrapf(err, "failed to create post")
		}
		if err := tx.Comment.WithContext(ctx).Create(&model.Comment{
			AuthorName:  t.T("Comment Robot"),
			AuthorEmail: "pub@gopub.tech",
			AuthorURL:   "https://code.gopub.tech/pub",
			PostID:      post.ID,
			Content:     t.T(`This is a sample comment.`),
		}); err != nil {
			return errors.Wrapf(err, "failed to create comment")
		}
		SetInstalled(ctx)
		SetTitle(ctx, req.SiteTitle)
		return nil
	})
}

// GetTitle 获取站点标题(缓存优先)
func GetTitle(ctx context.Context) string {
	val, _ := cache.GetOrFetch(ctx, caches.SiteTitle, func(ctx context.Context, key string) (any, error) {
		var s string
		o := query.Option
		option, err := o.WithContext(ctx).Where(o.Name.Eq(model.OptionNameSiteTitle)).First()
		if err == nil && option != nil {
			s = option.Value
		}
		return s, nil
	})
	return val.(string)
}

// 更新站点标题缓存
func SetTitle(ctx context.Context, s string) {
	cache.Set(ctx, caches.SiteTitle, s)
}
