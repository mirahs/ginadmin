package middleware

import (
	"ginadmin/common/menu"
	"ginadmin/conf"
	"ginadmin/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Auth 登录验证
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		urlFull := ctx.FullPath()
		// 如果是登录或登录出url直接放行
		if urlFull == conf.App.UrlLogin || urlFull == conf.App.UrlLogout {
			ctx.Next()
			return
		}

		// 没有登录跳转到登录url
		if !util.LoginCheck(ctx) {
			ctx.Redirect(http.StatusFound, conf.App.UrlLogin)
			ctx.Abort()
			return
		}

		// 权限验证(主页和访问拒绝url不需要验证权限)
		if urlFull != conf.App.UrlIndex && urlFull != conf.App.UrlDeny && !menu.Check(urlFull, util.GetType(ctx)) {
			ctx.Redirect(http.StatusFound, conf.App.UrlDeny)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
