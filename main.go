package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"path/filepath"

	"code.gopub.tech/logs"
	"code.gopub.tech/pub/settings"
	"code.gopub.tech/pub/webs"
	"github.com/gin-gonic/gin"
)

//go:embed lang/*.po
var lang embed.FS
var ctx = context.Background()
var dir = flag.String("dir", ".", "data dir")

func main() {
	MustInit()         // 初始化设置
	r := gin.Default() // 新建 gin 实例
	register(r)        // 注册路由
	// 启动
	addr := settings.AppConf.Addr
	logs.Info(ctx, "listen on %s", addr)
	logs.Info(ctx, "run: %+v", r.Run(addr))
}

func MustInit() {
	flag.Parse()
	abs, err := filepath.Abs(*dir)
	if err != nil {
		panic(err)
	}
	// logs
	logs.SetDefault(logs.NewLogger(logs.CombineHandlers(
		logs.NewHandler(), // console
		logs.NewHandler(
			logs.WithFile(filepath.Join(*dir, "logs", "app.log")),
		), // file
		// logs.NewHandler(logs.WithFile("logs/app.log.json"), logs.WithJSON()),
	)))
	logs.Info(ctx, "use data dir %q. starting app...", abs)
	if err := settings.ReadConfig(*dir); err != nil {
		panic(fmt.Sprintf("%+v", err))
	}
	webs.InitI18n(lang) // 尽早 load 翻译
}
