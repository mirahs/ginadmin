package admin

import (
	admin2 "ginadmin/service/admin"
	"github.com/flosch/pongo2/v4"
	"github.com/gin-gonic/gin"
	"net/http"
)


var sysService = admin2.NewSysService()


func SysPassword(ctx *gin.Context)  {

	if ctx.Request.Method == "POST" {
		err := sysService.Password(ctx)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"code": 1})
	} else {
		vmAdmUser := sysService.AccountInfo(ctx)

		ctx.HTML(http.StatusOK, "admin/sys/password.html", pongo2.Context{
			"account": vmAdmUser.Account,
		})
	}
}
