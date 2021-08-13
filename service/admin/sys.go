package admin

import (
	"errors"
	"ginadmin/repository"
	"ginadmin/util"
	"ginadmin/util/admin"
	"ginadmin/vm"
	"github.com/gin-gonic/gin"
)


type Sys struct {}


// 
func (*Sys) AccountInfo(ctx *gin.Context) vm.AdmUserVm {
	var vmAdmUser vm.AdmUserVm

	err := ctx.ShouldBind(&vmAdmUser)
	if err != nil || vmAdmUser.Account == "" {
		vmAdmUser.Account = admin.GetAccount(ctx)
	}

	return vmAdmUser
}

// 更改密码
func (*Sys) Password(ctx *gin.Context) (err error)  {
	var vmAdmUser vm.AdmUserVm
	err = ctx.ShouldBind(&vmAdmUser)
	if err != nil {
		err = errors.New("参数解析错误:" + err.Error())
		return
	}

	if vmAdmUser.Account == "" || vmAdmUser.Password == "" {
		err = errors.New("账号和新密码不能为空")
		return
	}

	repoAdmUser := repository.NewAdmUserRepository()
	repoAdmUser.UpdatePassword(vmAdmUser.Account, util.Md5(vmAdmUser.Password))

	return
}


func NewSysService() *Sys {
	return &Sys{}
}
