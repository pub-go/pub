package main

import (
	"context"
	"embed"

	"code.gopub.tech/logs"
	"code.gopub.tech/pub/webs"
	"github.com/gin-gonic/gin"
)

//go:embed lang
var lang embed.FS
var ctx = context.Background()

func main() {
	logs.Info(ctx, "Hello, World!") // 打个日志
	webs.InitI18n(lang)             // 尽早 load 翻译
	r := gin.Default()              // 新建 gin 实例
	register(r)                     // 注册路由
	r.Run()                         // 启动
}
