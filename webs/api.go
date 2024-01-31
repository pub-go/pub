package webs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Api(h func(ctx *gin.Context) (any, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result, err := h(ctx)
		if ctx.IsAborted() {
			return
		}
		ctx.JSON(http.StatusOK, Response(result, err))
	}
}

func Response(result any, err error) gin.H {
	if err != nil {
		return gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": result,
		}
	}
	return gin.H{
		"code": 0,
		"msg":  "ok",
		"data": result,
	}
}
