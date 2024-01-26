package webs

import (
	"context"
	"embed"
	"net/http"
	"os"
	"regexp"

	"code.gopub.tech/logs"
	"code.gopub.tech/pub/settings"
	"code.gopub.tech/tpl"
	"code.gopub.tech/tpl/html"
	"code.gopub.tech/tpl/types"
	"github.com/gin-gonic/gin"
)

// SetRender 设置模板渲染器
func SetRender(views embed.FS) gin.HandlerFunc {
	var (
		err      error
		viewExp  *regexp.Regexp
		viewPath = settings.AppConf.ViewPath
	)
	if viewPath != "" {
		pattern := settings.AppConf.ViewPattern
		if pattern == "" {
			pattern = "\\.html$"
		}
		viewExp, err = regexp.Compile(pattern)
		if err != nil {
			panic(err)
		}
	}
	render, err := tpl.NewHTMLRender(func(ctx context.Context) (types.TemplateManager, error) {
		m := html.NewTplManager()
		if viewPath != "" {
			logs.Info(ctx, "views path: %s, pattern: %s", viewPath, viewExp)
			// 使用 os.DirFS 实时读取文件夹
			return m, m.ParseWithRegexp(os.DirFS(viewPath), viewExp)
		}
		// 使用编译时嵌入的 embed.FS 资源
		logs.Info(ctx, "use internal views")
		return m, m.SetSubFS("views").ParseWithSuffix(views, ".html")
	}, tpl.WithHotReload(gin.IsDebugging()))
	if err != nil {
		panic(err)
	}
	return func(c *gin.Context) {
		c.Set("render", render)
		c.Next()
	}
}

func GetRender(ctx *gin.Context) types.ReloadableRender {
	return ctx.MustGet("render").(types.ReloadableRender)
}

// Render 渲染指定模板
func Render(ctx *gin.Context, name string, data gin.H) {
	logs.Info(ctx, "Render tpl=%s, data=%v", name, data)
	render := GetRender(ctx)
	data = WithI18n(ctx, data)
	tpl := render.Instance(ctx, name, data)
	ctx.Render(http.StatusOK, tpl)
}
