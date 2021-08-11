package admin

import (
	"ginadmin/config/menu"
	"ginadmin/model"
	"ginadmin/service/admin"
	"github.com/flosch/pongo2/v4"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)


func Index(ctx *gin.Context)  {
	user := ctx.MustGet("user").(model.AdmUser)

	ctx.HTML(http.StatusOK, "admin/index/index.html", pongo2.Context{
		"account":        user.Account,
		"user_type_name": "管理员",
		"menus":          menu.Menus(),
	})
}

func IndexLogin(ctx *gin.Context)  {
	if ctx.Request.Method == "POST" {
		indexService := admin.NewIndexService()

		err := indexService.Login(ctx)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"code": 1})
	} else {
		session := sessions.Default(ctx)
		if user := session.Get("user"); user != nil {
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
