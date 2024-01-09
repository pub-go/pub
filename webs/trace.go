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

// Trace 为每个请求设置一个唯一标记
func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := GetContext(c)
		trace := GenTraceID()
		c.Set("trace", trace)
		ctx = SetTraceID(ctx, trace)
		ctx = kv.Add(ctx,
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
		)
		SetContext(c, ctx)
		c.Next()
	}
}

func GenTraceID() string {
	now := time.Now()
	seq := atomic.AddUint64(&traceID, 1) % 10_000_000
	return fmt.Sprintf("%14s%09d%07d",
		now.Format("20060102150405"), now.Nanosecond(), seq)
}

func SetTraceID(ctx context.Context, s string) context.Context {
	ctx = kv.Add(ctx, "trace", s)
	return context.WithValue(ctx, &traceID, s)
}

func GetTraceID(ctx context.Context) string {
	v := ctx.Value(&traceID)
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}
