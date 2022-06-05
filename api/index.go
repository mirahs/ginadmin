package api

import (
	"ginadmin/common"
	"ginadmin/common/menu"
	"ginadmin/conf"
	"ginadmin/service"
	"ginadmin/util"
	"github.com/flosch/pongo2/v4"
	"github.com/gin-gonic/gin"
	"net/http"
)


// Index 主页
func Index(ctx *gin.Context)  {
	util.HTML(ctx, "index/index.html", pongo2.Context{
		"account":        util.GetAccount(ctx),
		"user_type_name": common.AdminUserTypesDesc[util.GetType(ctx)],
		"menus":          menu.Get(util.GetType(ctx)),
	})
}

// Login 登录
func Login(ctx *gin.Context)  {
	if ctx.Request.Method == "POST" {
		errMsg := service.Login(ctx)
		if errMsg != "" {
			ctx.JSON(http.StatusOK, gin.H{"code": 1, "msg": errMsg})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"code": 0})
	} else {
		if util.LoginCheck(ctx) {
			ctx.Redirect(http.StatusFound, conf.App.UrlIndex)
			return
		}
		util.HTML(ctx, "index/login.html", nil)
	}
}

// Logout 退出
func Logout(ctx *gin.Context) {
	service.Logout(ctx)

	ctx.Redirect(http.StatusFound, "login")
}

// Deny 访问拒绝
func Deny(ctx *gin.Context) {
	util.HTML(ctx, "index/deny.html", pongo2.Context{
		"account": util.GetAccount(ctx),
	})
}
