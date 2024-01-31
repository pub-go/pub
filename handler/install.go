package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type InstallReq struct {
	SiteTitle string `form:"site_title"`
	Username  string `form:"username"`
	Email     string `form:"email"`
	Password  string `form:"password"`
}

func Install(ctx *gin.Context) (any, error) {
	var req InstallReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		return nil, err
	}
	ctx.Redirect(http.StatusFound, "/login")
	ctx.Abort()
	return nil, nil
}
