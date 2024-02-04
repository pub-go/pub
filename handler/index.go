package handler

import (
	"code.gopub.tech/pub/dto"
	"code.gopub.tech/pub/service"
	"code.gopub.tech/pub/webs"
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	var req dto.QueryPostReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		showIndexPage(ctx, err)
		return
	}
	resp, err := service.QueryPosts(ctx, &req)
	showIndexPage(ctx, err, gin.H{"resp": resp})
}

func showIndexPage(ctx *gin.Context, err error, data ...gin.H) {
	webs.Render("index.html", append(data, gin.H{"err": err})...)(ctx)
}
