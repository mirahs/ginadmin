package middleware

import (
	"ginadmin/app/config"
	"ginadmin/app/config/menu"
	"ginadmin/app/util/admin"
	"github.com/gin-gonic/gin"
	"net/http"
)


// admin登录验证(如果是登录url直接放行, 其它要做登录验证, 如果没有登录就跳到登录url)
func AdminValidate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		urlFull := ctx.FullPath()
		if urlFull == config.AppInst.UrlLogin || urlFull == config.AppInst.UrlLogout {
			ctx.Next()
			return
		}

		if !admin.LoginCheck(ctx) {
			ctx.Redirect(http.StatusFound, config.AppInst.UrlLogin)
			ctx.Abort()
			return
		}

		if urlFull != config.AppInst.UrlIndex && urlFull != config.AppInst.UrlDeny && !menu.Check(urlFull, admin.GetType(ctx)) {
			ctx.Redirect(http.StatusFound, config.AppInst.UrlDeny)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
