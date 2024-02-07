package handler

import (
	"code.gopub.tech/pub/webs"
	"github.com/gin-gonic/gin"
)

func AdminPage(ctx *gin.Context) {
	webs.Render("admin/index.html")(ctx)
}
