package admin

import (
	"errors"
	"ginadmin/dto"
	"ginadmin/model"
	"ginadmin/repository"
	"ginadmin/util"
	"ginadmin/util/admin"
	"ginadmin/util/page"
	"ginadmin/vm"
	"github.com/gin-gonic/gin"
)


type Sys struct {}


func (*Sys) BindAdmUser(ctx *gin.Context) *vm.AdmUserVm {
	var vmAdmUser vm.AdmUserVm
	_ = ctx.ShouldBind(&vmAdmUser)
	return &vmAdmUser
}

func (sys *Sys) AccountInfo(ctx *gin.Context) *vm.AdmUserVm {
	vmAdmUser := sys.BindAdmUser(ctx)

	if vmAdmUser.Account == "" {
		vmAdmUser.Account = admin.GetAccount(ctx)
	}

	return vmAdmUser
}

func (*Sys) AdmUserVm2AdmUser(userVm *vm.AdmUserVm) *model.AdmUser {
	return &model.AdmUser{
		Id: userVm.Id,
		Account: userVm.Account,
		Type: userVm.Type,
		Remark: userVm.Remark,
	}
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
	repoAdmUser.UpdatePasswordByAccount(vmAdmUser.Account, util.Md5(vmAdmUser.Password))

	return
}

func (*Sys) MasterList(ctx *gin.Context) (*page.Info, []*dto.AdmUserDto) {
	var admUsers = make([]model.AdmUser, 0)
	pageInfo := page.Page(ctx, &admUsers)

	var admUserDtos = make([]*dto.AdmUserDto, 0)
	for _, admUser := range admUsers {
		admUserDtos = append(admUserDtos, dto.ToAdmUser(&admUser))
	}

	return pageInfo, admUserDtos
}


func NewSysService() *Sys {
	return &Sys{}
}
