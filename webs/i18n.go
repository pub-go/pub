package webs

import (
	"context"
	"embed"
	"strings"

	"code.gopub.tech/logs"
	"code.gopub.tech/logs/pkg/kv"
	"code.gopub.tech/pub/settings"
	"code.gopub.tech/tpl/exp"
	"github.com/gin-gonic/gin"
	"github.com/youthlin/t"
)

// InitI18n 初始化翻译
func InitI18n(defaultLangFs embed.FS) {
	if path := settings.AppConf.LangPath(); path != "" {
		t.Load(path)
		logs.Info(ctx, "set languange path: %v", path)
	} else {
		t.LoadFS(defaultLangFs)
		logs.Info(ctx, "use internal language translations")
	}
	t.SetLocale(settings.AppConf.Lang)
	logs.Info(ctx, "used locale: %s", t.UsedLocale())
}

// I18n 为每个请求决定使用哪种语言
func I18n(c *gin.Context) {
	ctx := GetContext(c)
	lang := t.GetUserLang(c.Request) // 获取浏览器偏好语言
	ctx = kv.Add(ctx, KeyLang, lang) // 在日志中打印
	ctx = t.SetCtxLocale(ctx, lang)  // 存在 ctx 里
	SetContext(c, ctx)               // 设置 ctx
	c.Next()
}

// WithI18n 将翻译工具函数等注入到模板变量中
func WithI18n(ctx context.Context, data gin.H) gin.H {
	t := t.WithContext(ctx) // 设置语言
	data["t"] = t
	data["__"] = t.T
	data["_x"] = t.X
	data["_n"] = func(msgID, msgIDPlural string, n any, args ...interface{}) string {
		return t.N64(msgID, msgIDPlural, exp.ToNumber[int64](n), args...)
	}
	data["_xn"] = func(msgCtx, msgID, msgIDPlural string, n any, args ...interface{}) string {
		return t.XN64(msgCtx, msgID, msgIDPlural, exp.ToNumber[int64](n), args...)
	}
	data[KeyLang] = strings.ReplaceAll(t.UsedLocale(), "_", "-")
	return data
}
