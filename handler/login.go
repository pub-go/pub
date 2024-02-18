package handler

import (
	"errors"
	"net/http"
	"strings"

	"code.gopub.tech/pub/dto"
	"code.gopub.tech/pub/service"
	"code.gopub.tech/pub/webs"
	"github.com/gin-gonic/gin"
)

func LoginPage(ctx *gin.Context) {
	showLoginPage(ctx, nil)
}

func showLoginPage(ctx *gin.Context, err error) {
	salt, err2 := service.GetStaticSalt(ctx)
	if err2 != nil {
		err = errors.Join(err, err2)
	}
	webs.Render("login.html", gin.H{"err": err, "salt": salt})(ctx)
}

func Login(ctx *gin.Context) {
	var req dto.LoginReq
	if err := ctx.ShouldBind(&req); err != nil {
		showLoginPage(ctx, err)
		return
	}
	if err := service.Login(ctx, &req); err != nil {
		showLoginPage(ctx, err)
		return
	}
	webs.SetUserInfo(ctx, &dto.LoginInfo{
		Username:  req.Username,
		UserAgent: ctx.Request.UserAgent(),
	})
	if strings.HasPrefix(req.Redirect, "/") {
		ctx.Redirect(http.StatusFound, req.Redirect)
	} else {
		ctx.Redirect(http.StatusFound, "/")
	}
}
