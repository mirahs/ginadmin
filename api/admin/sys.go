package admin

import (
	admin2 "ginadmin/service/admin"
	"github.com/flosch/pongo2/v4"
	"github.com/gin-gonic/gin"
	"net/http"
)


var sysService = admin2.NewSysService()


func SysPassword(ctx *gin.Context)  {
	vmAdmUser := sysService.PasswordInfo(ctx)
	if ctx.Request.Method == "POST" {

	} else {
		ctx.HTML(http.StatusOK, "admin/sys/password.html", pongo2.Context{
			"account": vmAdmUser.Account,
		})
	}
}
