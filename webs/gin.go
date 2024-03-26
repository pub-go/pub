package webs

import (
	"bytes"
	"context"
	"embed"
	"net/http"
	"os"
	"regexp"
	"sync/atomic"
	"time"

	"code.gopub.tech/errors"
	"code.gopub.tech/logs"
	"code.gopub.tech/pub/service"
	"code.gopub.tech/pub/settings"
	"code.gopub.tech/pub/util"
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
		viewPath = settings.AppConf.ViewPath()
	)
	if viewPath != "" {
		pattern := settings.AppConf.ViewPattern
		if pattern == "" {
			pattern = "\\.html$"
		}
		viewExp, err = regexp.Compile(pattern)
		if err != nil {
			util.Panic(ctx, errors.Wrapf(err, "invalid view pattern: %s", pattern))
		}
	}
	render, err := tpl.NewHTMLRender(func(ctx context.Context) (types.TemplateManager, error) {
		defer func(start time.Time) {
			d := time.Since(start)
			val := ctx.Value(KeyTplParseCost)
			if cost, ok := val.(*atomic.Int64); ok {
				cost.Store(int64(d))
			}
			logs.Info(ctx, "parse tpl cost: %v", d)
		}(time.Now())
		m := html.NewTplManager()
		if viewPath != "" {
			logs.Info(ctx, "views path: %s, pattern: %s", viewPath, viewExp)
			// 使用 os.DirFS 实时读取文件夹
			return m, m.ParseWithRegexp(os.DirFS(viewPath), viewExp)
		}
		// 使用编译时嵌入的 embed.FS 资源
		logs.Info(ctx, "use internal views")
		return m, m.SetSubFS("resource/views").ParseWithSuffix(views, ".html")
	}, tpl.WithHotReload(gin.IsDebugging()))
	if err != nil {
		util.Panic(ctx, err)
	}
	return func(c *gin.Context) {
		// 往上下文中注入 Render
		c.Set(KeyRender, render)
		c.Next()
	}
}

// GetRender 从上下文中拿出 Render (SetRender 时注入的)
func GetRender(ctx *gin.Context) types.ReloadableRender {
	return ctx.MustGet(KeyRender).(types.ReloadableRender)
}

// Render 渲染页面
func Render(tpl string, datas ...gin.H) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data = gin.H{}
		for _, m := range datas {
			for k, v := range m {
				data[k] = v
			}
		}
		render(ctx, tpl, data)
	}
}

// render 渲染指定模板
func render(ctx *gin.Context, name string, data gin.H) {
	data[KeyUser] = GetUser(ctx)
	reqStart := ctx.GetTime(KeyReqStart) // 请求开始时间
	serviceCost := time.Since(reqStart)  // 业务处理用时

	tplParseCost := &atomic.Int64{} // 模板解析用时
	parseCost := func() time.Duration {
		return time.Duration(tplParseCost.Load())
	}
	data[KeyTplParseCost] = parseCost
	// 如果 reload, 在 reload 时更新解析用时
	ctx.Set(KeyTplParseCost, tplParseCost)

	tplStart := time.Now()            // 页面渲染开始时间
	tplCost := func() time.Duration { // 页面渲染用时
		return time.Since(tplStart)
	}
	defer func() {
		var errs []error
		for _, e := range ctx.Errors {
			errs = append(errs, e)
		}
		logs.Info(ctx, "serviceCost=%v, Render tpl=%s, parseCost=%v, renderCost=%v err=%+v",
			serviceCost, name, parseCost(), tplCost(), errors.Join(errs...))
	}()

	data[KeyCtx] = ctx // 注入上下文 页面中可以 ctx.GetString
	data[KeyIsDebug] = gin.IsDebugging()
	data[keySiteTitle] = service.GetTitle(ctx)
	data[KeyServiceCost] = serviceCost // 请求用时

	data[KeyTplStart] = tplStart // 页面渲染开始时间
	data[KeyTplRenderCost] = tplCost
	data[KeyTotalCost] = func() time.Duration {
		return time.Since(reqStart) // 总体用时
	}

	data = WithI18n(ctx, data)           // 注入翻译函数
	data[KeySqlCount] = GetSqlCount(ctx) // 注入 sql 计数

	result, err := execute(ctx, name, data)
	if err != nil {
		cause := err
		data[KeyErr] = err
		if result, err = execute(ctx, "500.html", data); err != nil {
			err = errors.WithSecondary(cause, err)
		}
	}
	if err != nil {
		ctx.Data(http.StatusInternalServerError, "text/plain; charset=utf-8", []byte(errors.Detail(err)))
	} else {
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", result)
	}
}

func execute(ctx *gin.Context, name string, data any) ([]byte, error) {
	r := GetRender(ctx) // 获取 Render
	t, err := r.GetTemplate(ctx, name)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to GetTemplate: %s", name)
	}
	var sb bytes.Buffer
	err = t.Execute(&sb, data)
	if err != nil {
		logs.Error(ctx, "execute template %s error, data=%v, err=%+v", name, data, err)
		return nil, err
	}
	return sb.Bytes(), nil
}
