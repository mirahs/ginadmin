package middleware

import (
	"ginadmin/config/menu"
	"ginadmin/util/admin"
	"github.com/gin-gonic/gin"
	"net/http"
)


// admin登录验证(如果是登录url直接放行, 其它要做登录验证, 如果没有登录就跳到登录url)
func AdminValidate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		urlLogin := "/admin/index/login"
		urlLogout:= "/admin/index/logout"
		urlIndex := "/admin/index/index"
		urlNoAccess := "/admin/index/no_access"

		urlFull := ctx.FullPath()
		if urlFull == urlLogin || urlFull == urlLogout {
			ctx.Next()
			return
		}

		if !admin.LoginCheck(ctx) {
			ctx.Redirect(http.StatusFound, urlLogin)
			ctx.Abort()
			return
		}

		if urlFull != urlIndex && urlFull != urlNoAccess && !menu.Check(urlFull, admin.GetAccountType(ctx)) {
			ctx.Redirect(http.StatusFound, urlNoAccess)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
