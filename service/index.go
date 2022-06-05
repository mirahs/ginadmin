package service

import (
	"ginadmin/model"
	"ginadmin/util"
	"ginadmin/vm"
	"ginadmin/vo"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)


func Login(ctx *gin.Context) string  {
	var voAdmUser vo.AdmUserLogin

	err := util.GinBindValid(ctx, &voAdmUser)
	if err != nil {
		return "参数错误"
	}

	vmAU := vm.AdminUser{Account: voAdmUser.Account, Password: voAdmUser.Password}

	ip := ctx.ClientIP()
	address := util.IpAddress(ip)

	admUser, err := vmAU.GetByAccount()
	if err != nil {
		return "DB错误"
	}
	if admUser.Id == 0 {
		return "账号不存在"
	}

	passwordMd5 := util.Md5(voAdmUser.Password)
	if passwordMd5 != admUser.Password {
		model.LogAdmUserLoginAddFailed(admUser.Id, ip, address, "密码错误")
		return "密码错误"
	}

	loginSessionSet(ctx, admUser)

	model.AULoginAfter(admUser, ip)
	model.LogAdmUserLoginAddSuccess(admUser.Id, ip, address, "")

	return ""
}

func Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	_ = session.Save()
}
