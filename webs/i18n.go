package webs

import (
	"embed"
	"strings"

	"code.gopub.tech/logs/pkg/kv"
	"code.gopub.tech/tpl/exp"
	"github.com/gin-gonic/gin"
	"github.com/youthlin/t"
)

// InitI18n 初始化翻译
func InitI18n(lang embed.FS) {
	if gin.IsDebugging() {
		t.Load("./lang")
	} else {
		t.LoadFS(lang)
	}
}

// I18n 为每个请求决定使用哪种语言
func I18n() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := GetContext(c)
		lang := t.GetUserLang(c.Request) // 获取浏览器偏好语言
		ctx = kv.Add(ctx, "lang", lang)  // 在日志中打印
		ctx = t.SetCtxLocale(ctx, lang)  // 存在 ctx 里
		SetContext(c, ctx)               // 设置 ctx
		c.Next()
	}
}

// WithI18n 将翻译工具函数等注入到模板变量中
func WithI18n(c *gin.Context, data gin.H) gin.H {
	if data == nil {
		data = make(gin.H)
	}
	data["ctx"] = c
	ctx := GetContext(c)    // 取出 ctx
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
	data["lang"] = strings.ReplaceAll(t.UsedLocale(), "_", "-")
	return data
}
