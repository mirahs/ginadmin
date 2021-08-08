package middleware

import (
	"github.com/gin-contrib/sessions"
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

		session := sessions.Default(ctx)
		if isLogin := session.Get("isLogin"); isLogin == nil || !isLogin.(bool) {
			ctx.Redirect(http.StatusFound, urlLogin)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
