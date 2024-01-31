package service

import (
	"context"

	"code.gopub.tech/errors"
	"code.gopub.tech/pub/dal/model"
	"code.gopub.tech/pub/dal/query"
	"gorm.io/gorm"
)

var installed *bool

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
	*installed = true
}
