package main

import (
	"embed"
	"net/http"

	"code.gopub.tech/pub/handler"
	"code.gopub.tech/pub/webs"
	"github.com/gin-gonic/gin"
)

//go:embed views
var views embed.FS

func register(r *gin.Engine) {
	// 开启后 gin.Context 才能回退到 c.Request.Context()
	// 就不需要先 ctx := webs.GetContext( ginCtx ) 再使用 ctx 而可以直接传 ginCtx 了
	r.ContextWithFallback = true

	// 设置好各中间件
	r.Use(webs.Trace(), webs.I18n(), webs.SetRender(views), webs.Install())

	// 路由逻辑
	front(r)
	install(r)
	admin(r)
}

func front(g gin.IRouter) {
	g.GET("/", webs.Render("index.html"))
}

func install(g gin.IRouter) {
	g.GET("/install/", webs.Render("install.html"))
	g.POST("/install/", webs.Api(handler.Install))
}

func admin(g gin.IRouter) {
	g.GET("/admin/render/reload", func(ctx *gin.Context) {
		err := webs.GetRender(ctx).Reload(ctx)
		ctx.JSON(http.StatusOK, gin.H{
			"err": err,
		})
	})
}
