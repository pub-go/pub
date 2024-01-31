package webs

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"code.gopub.tech/logs/pkg/kv"
	"github.com/gin-gonic/gin"
)

var traceID uint64
var sqlCount struct{}

const (
	KeyReqStart    = "reqStart"
	KeyTrace       = "trace"
	KeyLang        = "lang"
	KeySqlCount    = "sqlCount"
	KeyRender      = "render"
	KeyCtx         = "ctx"
	KeyServiceCost = "serviceCost"
	KeyTplStart    = "tplStart"
	KeyTplCost     = "tplCost"
	KeyTotalCost   = "totalCost"
	HeaderTrace    = "X-Trace-ID"
)

// Trace 为每个请求设置一个唯一标记
func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		ctx = InitSqlCount(ctx)

		SetContext(c, ctx)
		c.Next()
	}
}

// GenTraceID 生成 traceID
func GenTraceID() string {
	now := time.Now()
	seq := atomic.AddUint64(&traceID, 1) % 10_000_000
	return fmt.Sprintf("%14s%09d%07d",
		now.Format("20060102150405"), now.Nanosecond(), seq)
}

// InitSqlCount 初始化 sql 计数器
func InitSqlCount(ctx context.Context) context.Context {
	var count atomic.Uint32
	return context.WithValue(ctx, sqlCount, &count)
}

// AddSqlCount sql 计数加一
func AddSqlCount(ctx context.Context) {
	value := ctx.Value(sqlCount)
	if count, ok := value.(*atomic.Int32); ok {
		count.Add(1)
	}
}

// WithSqlCount 往渲染数据中添加 sql 计数
func WithSqlCount(ctx context.Context, data gin.H) gin.H {
	var count int64
	value := ctx.Value(sqlCount)
	if i, ok := value.(*atomic.Int32); ok {
		count = int64(i.Load())
	}
	data[KeySqlCount] = count
	return data
}
