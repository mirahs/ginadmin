package middleware

import (
	"ginadmin/common/menu"
	"ginadmin/config"
	"ginadmin/util/admin"
	"github.com/gin-gonic/gin"
	"net/http"
)


// 登录验证
func AdminValidate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		urlFull := ctx.FullPath()
		// 如果是登录或登录出url直接放行
		if urlFull == config.App.UrlLogin || urlFull == config.App.UrlLogout {
			ctx.Next()
			return
		}

		// 没有登录跳转到登录url
		if !admin.LoginCheck(ctx) {
			ctx.Redirect(http.StatusFound, config.App.UrlLogin)
			ctx.Abort()
			return
		}

		// 权限验证(主页和访问拒绝url不需要验证权限)
		if urlFull != config.App.UrlIndex && urlFull != config.App.UrlDeny && !menu.Check(urlFull, admin.GetType(ctx)) {
			ctx.Redirect(http.StatusFound, config.App.UrlDeny)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
