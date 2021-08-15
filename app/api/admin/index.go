package admin

import (
	"ginadmin/app/config"
	"ginadmin/app/config/menu"
	"ginadmin/app/service/admin"
	util_admin "ginadmin/app/util/admin"
	"github.com/flosch/pongo2/v4"
	"github.com/gin-gonic/gin"
	"net/http"
)


var serviceIndex = admin.NewServiceIndex()


// 后台主页
func Index(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK, "admin/index/index.html", pongo2.Context{
		"account":        util_admin.GetAccount(ctx),
		"user_type_name": config.GetTypeName(util_admin.GetType(ctx)),
		"menus":          menu.Get(util_admin.GetType(ctx)),
	})
}

// 后台登录
func IndexLogin(ctx *gin.Context)  {
	if ctx.Request.Method == "POST" {
		err := serviceIndex.Login(ctx)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"code": 1})
	} else {
		if util_admin.LoginCheck(ctx) {
			ctx.Redirect(http.StatusFound, "index")
			return
		}
		ctx.HTML(http.StatusOK, "admin/index/login.html", pongo2.Context{})
	}
}

// 后台退出
func IndexLogout(ctx *gin.Context) {
	serviceIndex.Logout(ctx)
	ctx.Redirect(http.StatusFound, "login")
}

// 访问拒绝
func IndexDeny(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/index/deny.html", pongo2.Context{
		"account": util_admin.GetAccount(ctx),
	})
}
