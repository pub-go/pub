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

func SetInstalled(ctx context.Context) {
	cache.Set(ctx, caches.Installed, true)
}

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

func Install(ctx context.Context, req *dto.InstallReq) error {
	salt, err := GetStaticSalt(ctx)
	if err != nil {
		return err
	}
	if req.Salt != salt {
		return errors.Errorf("invalid request, mismatch salt")
	}
	return query.Q.Transaction(func(tx *query.Query) error {
		u := tx.User
		salt := util.RandStr(16) // 每个用户在后端再单独随机加盐
		pass := sha512.Sum512([]byte(req.Password + salt))
		user := &model.User{
			Username: req.Username,
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
		SetInstalled(ctx)
		SetTitle(ctx, req.SiteTitle)
		return nil
	})
}

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

func SetTitle(ctx context.Context, s string) {
	cache.Set(ctx, caches.SiteTitle, s)
}
