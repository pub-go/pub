package service

import (
	"context"
	"crypto/sha512"
	"fmt"

	"code.gopub.tech/errors"
	"code.gopub.tech/pub/dal/query"
	"code.gopub.tech/pub/reqs"
)

var ErrLoginFailed = fmt.Errorf("wrong username or password")

func Login(ctx context.Context, req *reqs.LoginReq) error {
	u := query.User
	user, err := u.WithContext(ctx).Where(u.Username.Eq(req.Username)).First()
	if err != nil {
		return errors.WithSecondary(ErrLoginFailed, err)
	}
	pass := sha512.Sum512([]byte(req.Password + user.Salt))
	if string(user.Password) != string(pass[:]) {
		return errors.WithSecondary(ErrLoginFailed, err)
	}
	return nil
}
