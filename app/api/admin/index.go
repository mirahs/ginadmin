package admin

import (
	"ginadmin/app/config"
	"ginadmin/app/config/menu"
	"ginadmin/app/service/admin"
	admin2 "ginadmin/app/util/admin"
	"github.com/flosch/pongo2/v4"
	"github.com/gin-gonic/gin"
	"net/http"
)


var indexService = admin.NewIndexService()


// 后台主页
func Index(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK, "admin/index/index.html", pongo2.Context{
		"account":        admin2.GetAccount(ctx),
		"user_type_name": config.GetTypeName(admin2.GetType(ctx)),
		"menus":          menu.Get(admin2.GetType(ctx)),
	})
}

// 后台登录
func IndexLogin(ctx *gin.Context)  {
	if ctx.Request.Method == "POST" {
		err := indexService.Login(ctx)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"code": 1})
	} else {
		if admin2.LoginCheck(ctx) {
			ctx.Redirect(http.StatusFound, "index")
			return
		}
		ctx.HTML(http.StatusOK, "admin/index/login.html", pongo2.Context{})
	}
}

// 后台退出
func IndexLogout(ctx *gin.Context) {
	indexService.Logout(ctx)
	ctx.Redirect(http.StatusFound, "login")
}

// 访问拒绝
func IndexDeny(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/index/deny.html", pongo2.Context{
		"account": admin2.GetAccount(ctx),
	})
}
