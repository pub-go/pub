package handler

import (
	"net/http"

	"code.gopub.tech/errors"
	"code.gopub.tech/pub/dto"
	"code.gopub.tech/pub/service"
	"code.gopub.tech/pub/webs"
	"github.com/gin-gonic/gin"
)

// InstallPage 安装页面
func InstallPage(ctx *gin.Context) {
	if checkInstalled(ctx) {
		return
	}
	showInstallPage(ctx, nil)
}

// checkInstalled 如果已经安装跳转到首页
func checkInstalled(ctx *gin.Context) bool {
	if service.Installed(ctx) {
		ctx.Redirect(http.StatusFound, "/")
		return true
	}
	return false
}

// showInstallPage 渲染页面
func showInstallPage(ctx *gin.Context, err error) {
	salt, err2 := service.GetStaticSalt(ctx)
	if err2 != nil {
		err = errors.Join(err, err2)
	}
	webs.Render("install.html", gin.H{"err": err, "salt": salt})(ctx)
}

// Install 执行安装动作
func Install(ctx *gin.Context) {
	if checkInstalled(ctx) {
		return
	}
	var req dto.InstallReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		showInstallPage(ctx, err)
		return
	}
	if err := service.Install(ctx, &req); err != nil {
		showInstallPage(ctx, err)
		return
	}
	ctx.Redirect(http.StatusFound, "/login")
}
