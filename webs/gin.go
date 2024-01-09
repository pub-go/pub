package webs

import (
	"embed"
	"net/http"
	"os"

	"code.gopub.tech/logs"
	"code.gopub.tech/tpl"
	"code.gopub.tech/tpl/html"
	"code.gopub.tech/tpl/types"
	"github.com/gin-gonic/gin"
)

// SetRender 设置模板渲染器
func SetRender(views embed.FS) gin.HandlerFunc {
	var hotReload = gin.IsDebugging()
	render, err := tpl.NewHTMLRender(func() (types.TemplateManager, error) {
		m := html.NewTplManager()
		if hotReload {
			// 使用 os.DirFS 实时读取文件夹
			return m, m.ParseWithSuffix(os.DirFS("views"), ".html")
		}
		// 使用编译时嵌入的 embed.FS 资源
		return m, m.SetSubFS("views").ParseWithSuffix(views, ".html")
	}, hotReload)
	if err != nil {
		panic(err)
	}
	return func(c *gin.Context) {
		c.Set("render", render)
		c.Next()
	}
}

func GetRender(c *gin.Context) types.HTMLRender {
	return c.MustGet("render").(types.HTMLRender)
}

// Render 渲染指定模板
func Render(c *gin.Context, name string, data gin.H) {
	ctx := GetContext(c)
	c.Header("X-Trace-ID", GetTraceID(ctx))
	logs.Info(ctx, "Render tpl=%s, data=%v", name, data)
	render := GetRender(c)
	data = WithI18n(c, data)
	tpl := render.Instance(name, data)
	c.Render(http.StatusOK, tpl)
}
