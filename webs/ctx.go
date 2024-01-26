package webs

import (
	"context"

	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

func SetContext(c *gin.Context, ctx context.Context) {
	c.Request = c.Request.WithContext(ctx)
}

func GetContext(c *gin.Context) context.Context {
	return c.Request.Context()
}
