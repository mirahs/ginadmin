package admin

import (
	"encoding/json"
	"errors"
	"ginadmin/repository"
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

	repoAdmUser := repository.NewAdmUserRepository()

	admUser := repoAdmUser.GetByAccount(vmAdmUser.Account)
	if admUser.ID == 0 {
		err = errors.New("账号不存在")
		return
	}

	passwordMd5 := util.Md5(vmAdmUser.Password)
	if passwordMd5 != admUser.Password {
		err = errors.New("密码错误")
		return
	}

	// session不能序列化struct(不确定), 只能转换成json字符串
	admUserJsonBytes, _ := json.Marshal(admUser)

	session := sessions.Default(ctx)
	session.Set("user", string(admUserJsonBytes))
	err = session.Save()

	return
}


func NewIndexService() *Index {
	return &Index{}
}
