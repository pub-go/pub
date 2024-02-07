package webs

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"code.gopub.tech/logs/pkg/kv"
	"code.gopub.tech/pub/util"
	"github.com/gin-gonic/gin"
)

const (
	KeyUser          = "user"
	KeyReqStart      = "reqStart"
	KeyTrace         = "trace"
	KeyLang          = "lang"
	keySiteTitle     = "siteTitle"
	KeySqlCount      = "sqlCount"
	KeyRender        = "render"
	KeyCtx           = "ctx"
	KeyIsDebug       = "isDebug"
	KeyServiceCost   = "serviceCost"
	KeyTplStart      = "tplStart"
	KeyTplParseCost  = "tplParseCost"
	KeyTplRenderCost = "tplRenderCost"
	KeyTotalCost     = "totalCost"
	KeyErr           = "err"
	HeaderTrace      = "X-Trace-ID"
)

// Trace 为每个请求设置一个唯一标记
func Trace(c *gin.Context) {

	// 请求开始时间
	now := time.Now()
	c.Set(KeyReqStart, now)

	// 生成一个唯一标记
	trace := GenTraceID()
	c.Set(KeyTrace, trace)
	c.Header(HeaderTrace, trace)

	ctx := GetContext(c)
	ctx = kv.Add(ctx,
		"method", c.Request.Method,
		"path", c.Request.URL.Path,
		KeyTrace, trace,
	)

	// SQL 查询计数
	c.Set(KeySqlCount, &atomic.Int64{})

	SetContext(c, ctx)
	c.Next()

}

// GenTraceID 生成 traceID
func GenTraceID() string {
	now := time.Now()
	return fmt.Sprintf("%14s%09d%07s",
		now.Format("20060102150405"), now.Nanosecond(), util.RandStr(7))
}

// AddSqlCount sql 计数加一
func AddSqlCount(ctx context.Context) {
	value := ctx.Value(KeySqlCount)
	if count, ok := value.(*atomic.Int64); ok {
		count.Add(1)
	}
}

func GetSqlCount(ctx *gin.Context) int64 {
	value, ok := ctx.Get(KeySqlCount)
	if !ok {
		return 0
	}
	if count, ok := value.(*atomic.Int64); ok {
		return count.Load()
	}
	return 0
}
