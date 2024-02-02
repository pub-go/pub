package service

import (
	"context"
	"crypto/sha512"

	"code.gopub.tech/errors"
	"code.gopub.tech/pub/dal/model"
	"code.gopub.tech/pub/dal/query"
	"code.gopub.tech/pub/reqs"
	"code.gopub.tech/pub/util"
	"gorm.io/gorm"
)

var installed *bool
var title *string
var salt string

// Installed 是否已经安装
func Installed(ctx context.Context) bool {
	if installed == nil {
		b := queryInstalled(ctx)
		installed = &b
	}
	return *installed
}

func queryInstalled(ctx context.Context) bool {
	o := query.Option
	do := o.WithContext(ctx)
	option, err := do.Where(o.Name.Eq(model.OptionNameInstalled)).
		Limit(1).
		First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return option.Value == model.OptionValueYes
}

func SetInstalled() {
	b := true
	installed = &b
}

func GetStaticSalt(ctx context.Context) (string, error) {
	if salt == "" {
		o := query.Option
		option, err := o.WithContext(ctx).
			Attrs(o.Value.Value(util.RandStr(16))).
			Where(o.Name.Eq(model.OptionNameSalt)).
			FirstOrCreate()
		if err != nil {
			return "", errors.Wrapf(err, "failed to get salt")
		}
		salt = option.Value
	}
	return salt, nil
}

func Install(ctx context.Context, req *reqs.InstallReq) error {
	if req.Salt != salt {
		return errors.Errorf("invalid request, mismatch salt")
	}
	return query.Q.Transaction(func(tx *query.Query) error {
		u := tx.User
		salt := util.RandStr(16)
		pass := sha512.Sum512([]byte(req.Password + salt))
		if err := u.WithContext(ctx).Create(&model.User{
			Username: req.Username,
			Email:    req.Email,
			Password: pass[:],
			Salt:     salt,
		}); err != nil {
			return err
		}
		o := tx.Option
		option := &model.Option{Name: model.OptionNameSiteTitle, Value: req.SiteTitle}
		if err := o.WithContext(ctx).Save(option); err != nil {
			return err
		}
		option = &model.Option{Name: model.OptionNameInstalled, Value: model.OptionValueYes}
		if err := o.WithContext(ctx).Save(option); err != nil {
			return err
		}
		SetInstalled()
		SetTitle(req.SiteTitle)
		return nil
	})
}

func GetTitle(ctx context.Context) string {
	if title == nil {
		var s string
		o := query.Option
		option, err := o.WithContext(ctx).Where(o.Name.Eq(model.OptionNameSiteTitle)).First()
		if err == nil && option != nil {
			s = option.Value
		}
		title = &s
	}
	return *title
}
func SetTitle(s string) {
	title = &s
}
