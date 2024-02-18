package webs

import (
	"net/http"

	"code.gopub.tech/pub/dal/model"
	"github.com/gin-gonic/gin"
)

func Admin(ctx *gin.Context) {
	user := GetUser(ctx)
	if user != nil {
		if user.HasRole(model.RoleSuperAdmin) {
			ctx.Next()
			return
		}
	}
	ctx.Redirect(http.StatusFound, "/login?redirect=/admin/")
	ctx.Abort()
}
