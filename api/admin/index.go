package admin

import (
	"ginadmin/menu"
	"github.com/flosch/pongo2/v4"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)


func Index(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK, "admin/index/index.html", pongo2.Context{
		"account": "mirahs",
		"user_type_name": "管理员",
		"menus": menu.Menus(),
	})
}

func IndexLogin(ctx *gin.Context)  {
	if ctx.Request.Method == "POST" {
		account := ctx.DefaultPostForm("account", "")
		password:= ctx.DefaultPostForm("password", "")
		if account == "admin" && password == "admin" {
			session := sessions.Default(ctx)
			session.Set("isLogin", true)
			session.Save()

			ctx.JSON(http.StatusOK, gin.H{"code": 1})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"code": 0, "msg": "账号或密码错误"})
		}
	} else {
		session := sessions.Default(ctx)
		if isLogin := session.Get("isLogin"); isLogin != nil && isLogin.(bool) {
			ctx.Redirect(http.StatusFound, "index")
			return
		}
		ctx.HTML(http.StatusOK, "admin/index/login.html", pongo2.Context{})
	}
}

func IndexLogout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()

	ctx.Redirect(http.StatusFound, "login")
}
