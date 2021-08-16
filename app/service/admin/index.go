package admin

import (
	"errors"
	"ginadmin/app/thirdparty"
	"ginadmin/app/util"
	"ginadmin/app/util/admin"
	"ginadmin/app/vm"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)


type Index struct {
	base
}


func (index *Index) Login(ctx *gin.Context) (err error)  {
	var vmAdmUser vm.AdmUserVm
	err = ctx.ShouldBind(&vmAdmUser)
	if err != nil {
		err = errors.New("参数解析错误:" + err.Error())
		return
	}

	if vmAdmUser.Account == "" || vmAdmUser.Password == "" {
		err = errors.New("账号和密码不能为空")
		return
	}

	ip := ctx.ClientIP()
	address := thirdparty.IpAddress(ip)

	var remark string

	admUser := index.RepoAdmUser.GetByAccount(vmAdmUser.Account)
	if admUser.Id == 0 {
		remark = "账号不存在"
		err = errors.New(remark)
		index.RepoLogAdmUserLogin.AddFailed(vmAdmUser.Account, ip, address, remark)
		return
	}

	passwordMd5 := util.Md5(vmAdmUser.Password)
	if passwordMd5 != admUser.Password {
		remark = "密码错误"
		err = errors.New(remark)
		index.RepoLogAdmUserLogin.AddFailed(vmAdmUser.Account, ip, address, remark)
		return
	}

	admin.LoginSessionSet(ctx, admUser)

	index.RepoAdmUser.LoginUpdate(admUser, ip)
	index.RepoLogAdmUserLogin.AddSuccess(vmAdmUser.Account, ip, address, remark)

	return
}

func (*Index) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	_ = session.Save()
}
