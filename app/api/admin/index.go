package admin

import (
	"ginadmin/app/config"
	"ginadmin/app/config/menu"
	"ginadmin/app/util/admin"
	"github.com/flosch/pongo2/v4"
	"github.com/gin-gonic/gin"
	"net/http"
)


// 主页
func Index(ctx *gin.Context)  {
	admin.HTML(ctx, "admin/index/index.html", pongo2.Context{
		"account":        admin.GetAccount(ctx),
		"user_type_name": config.GetTypeName(admin.GetType(ctx)),
		"menus":          menu.Get(admin.GetType(ctx)),
	})
}

// 登录
func IndexLogin(ctx *gin.Context)  {
	if ctx.Request.Method == "POST" {
		err := serviceIndex.Login(ctx)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"code": 1})
	} else {
		if admin.LoginCheck(ctx) {
			ctx.Redirect(http.StatusFound, "index")
			return
		}
		admin.HTML(ctx, "admin/index/login.html", nil)
	}
}

// 退出
func IndexLogout(ctx *gin.Context) {
	serviceIndex.Logout(ctx)
	ctx.Redirect(http.StatusFound, "login")
}

// 访问拒绝
func IndexDeny(ctx *gin.Context) {
	admin.HTML(ctx, "admin/index/deny.html", pongo2.Context{
		"account": admin.GetAccount(ctx),
	})
}
