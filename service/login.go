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

var errLoginFailed = fmt.Errorf(t.Mark.T("wrong username or password"))

// Login 执行登录检查
func Login(ctx context.Context, req *dto.LoginReq) error {
	u := query.User
	// 先找到用户名
	user, err := u.WithContext(ctx).Where(u.Username.Eq(req.Username)).First()
	if err != nil {
		return errors.WithSecondary(ErrLoginFailed(ctx), err)
	}
	// 再获取用户的 salt 并对输入的密码加盐
	pass := sha512.Sum512([]byte(req.Password + user.Salt))
	// 然后比较密码
	if string(user.Password) != string(pass[:]) {
		return errors.WithSecondary(ErrLoginFailed(ctx), err)
	}
	return nil // 登录成功
}

func ErrLoginFailed(ctx context.Context) error {
	t := t.WithContext(ctx)
	err := errors.Errorf(t.T(errLoginFailed.Error()))
	if err.Error() != errLoginFailed.Error() {
		// 翻译+原文
		return errors.WithSecondary(err, errLoginFailed)
	}
	return err
}
