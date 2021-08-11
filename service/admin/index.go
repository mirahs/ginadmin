package admin

import (
	"encoding/json"
	"errors"
	"ginadmin/repository"
	"ginadmin/thirdparty/ip2region"
	"ginadmin/util"
	"ginadmin/vm"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)


type Index struct {}


func (i *Index) Login(ctx *gin.Context) (err error)  {
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
	address := ip2region.GetAddress(ip)

	var remark string
	repoAdmUser := repository.NewAdmUserRepository()
	repoLogAdmUserLogin := repository.NewLogAdmUserLoginRepository()

	admUser := repoAdmUser.GetByAccount(vmAdmUser.Account)
	if admUser.ID == 0 {
		remark = "账号不存在"
		err = errors.New(remark)
		repoLogAdmUserLogin.AddFailed(vmAdmUser.Account, ip, address, remark)
		return
	}

	passwordMd5 := util.Md5(vmAdmUser.Password)
	if passwordMd5 != admUser.Password {
		remark = "密码错误"
		err = errors.New(remark)
		repoLogAdmUserLogin.AddFailed(vmAdmUser.Account, ip, address, remark)
		return
	}

	// session不能序列化struct(不确定), 只能转换成json字符串
	admUserJsonBytes, _ := json.Marshal(admUser)

	session := sessions.Default(ctx)
	session.Set("user", string(admUserJsonBytes))
	err = session.Save()

	repoAdmUser.LoginUpdate(admUser, ip)
	repoLogAdmUserLogin.AddSuccess(vmAdmUser.Account, ip, address, remark)

	return
}


func NewIndexService() *Index {
	return &Index{}
}
