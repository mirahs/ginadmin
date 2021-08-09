package admin

import (
	"encoding/json"
	"ginadmin/config/menu"
	"ginadmin/model"
	"ginadmin/util"
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
		account := ctx.DefaultPostForm("account", "")
		password:= ctx.DefaultPostForm("password", "")

		if account == "" || password == "" {
			ctx.JSON(http.StatusOK, gin.H{"code": 0, "msg": "账号和密码不能为空"})
			return
		}

		var admUser = model.AdmUser{}
		model.Db.Find(&admUser, "`account`=?", account)
		if admUser.ID == 0 {
			ctx.JSON(http.StatusOK, gin.H{"code": 0, "msg": "账号不存在"})
			return
		}

		passwordMd5 := util.Md5(password)
		if passwordMd5 != admUser.Password {
			ctx.JSON(http.StatusOK, gin.H{"code": 0, "msg": "密码错误"})
			return
		}

		// session不能序列化struct(不认定), 只能转换成json字符串
		admUserJsonBytes, _ := json.Marshal(admUser)

		session := sessions.Default(ctx)
		session.Set("user", string(admUserJsonBytes))
		session.Save()

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
