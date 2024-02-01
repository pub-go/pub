package main

import (
	"context"
	"embed"
	"flag"
	"path/filepath"

	"code.gopub.tech/errors"
	"code.gopub.tech/logs"
	"code.gopub.tech/pub/dal"
	"code.gopub.tech/pub/settings"
	"code.gopub.tech/pub/util"
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
	flag.Parse() // 支持指定数据目录
	dir := *dir  // 默认当前文件夹 转为绝对路径
	abs, err := filepath.Abs(dir)
	if err != nil {
		util.Panic(ctx, errors.Wrapf(err, "failed to get abs path of %s", dir))
	}
	// logs 日志输出控制台、文件
	logs.SetDefault(logs.NewLogger(logs.CombineHandlers(
		logs.NewHandler(), // console
		logs.NewHandler(logs.WithFile(filepath.Join(abs, "logs", "app.log"))), // file
		// logs.NewHandler(logs.WithFile("logs/app.log.json"), logs.WithJSON()),  // json
	)))
	logs.Info(ctx, "use data dir %q. starting app...", abs)
	// 读取配置文件
	if err := settings.ReadConfig(abs); err != nil {
		util.Panic(ctx, errors.Wrapf(err, "ReadConfig"))
	}
	webs.InitI18n(lang) // 尽早 load 翻译
	dal.MustInit(abs)   // 连接数据库
}
