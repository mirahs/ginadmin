package middleware

import (
	"ginadmin/util/admin"
	"github.com/gin-gonic/gin"
	"net/http"
)


// admin登录验证(如果是登录url直接放行, 其它要做登录验证, 如果没有登录就跳到登录url)
func AdminValidate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		urlLogin := "/admin/index/login"
		urlFull := ctx.FullPath()
		if urlFull == urlLogin {
			ctx.Next()
			return
		}

		if !admin.LoginCheck(ctx) {
			ctx.Redirect(http.StatusFound, urlLogin)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
