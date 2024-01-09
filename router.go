package main

import (
	"embed"

	"code.gopub.tech/pub/webs"
	"github.com/gin-gonic/gin"
)

//go:embed views
var views embed.FS

func register(r *gin.Engine) {
	// 设置好各中间件
	r.Use(webs.Trace(), webs.I18n(), webs.SetRender(views))
	// 路由逻辑
	r.GET("/", func(c *gin.Context) {
		webs.Render(c, "index.html", nil) // 执行模板
	})
}
