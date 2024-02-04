package service

import (
	"context"
	"crypto/sha512"
	"fmt"

	"code.gopub.tech/errors"
	"code.gopub.tech/pub/dal/query"
	"code.gopub.tech/pub/dto"
	"github.com/youthlin/t"
)

var errLoginFailed = fmt.Errorf(t.T("wrong username or password"))

func Login(ctx context.Context, req *dto.LoginReq) error {
	u := query.User
	user, err := u.WithContext(ctx).Where(u.Username.Eq(req.Username)).First()
	if err != nil {
		return errors.WithSecondary(ErrLoginFailed(ctx), err)
	}
	pass := sha512.Sum512([]byte(req.Password + user.Salt))
	if string(user.Password) != string(pass[:]) {
		return errors.WithSecondary(ErrLoginFailed(ctx), err)
	}
	return nil
}

func ErrLoginFailed(ctx context.Context) error {
	t := t.WithContext(ctx)
	err := errors.Errorf(t.T(errLoginFailed.Error()))
	return errors.WithSecondary(err, errLoginFailed)
}
