package webs

import (
	"net/http"
	"strings"

	"code.gopub.tech/pub/service"
	"github.com/gin-gonic/gin"
)

// Install 未安装拦截
func Install() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if strings.HasPrefix(path, "/install/") || strings.HasPrefix(path, "/static/") || path == "/favicon.ico" {
			ctx.Next() // 安装页面、静态资源放过
			return
		}
		if service.Installed(ctx) {
			ctx.Next() // 已安装放过
			return
		}
		// 跳转安装页面
		ctx.Redirect(http.StatusFound, "/install/")
		ctx.Abort() // 重要
	}
}
