package admin

import (
	"errors"
	"ginadmin/app/dto"
	"ginadmin/app/model"
	"ginadmin/app/util"
	"ginadmin/app/util/admin"
	"ginadmin/app/util/page"
	"ginadmin/app/vm"
	"github.com/gin-gonic/gin"
)


type Sys struct{
	base
}


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
		Id:      userVm.Id,
		Account: userVm.Account,
		Type:    userVm.Type,
		Remark:  userVm.Remark,
	}
}

// 更改密码
func (sys *Sys) Password(ctx *gin.Context) (err error) {
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

	admUser := sys.RepoAdmUser.GetByAccount(vmAdmUser.Account)
	if admUser.Id == 0 {
		err = errors.New("账号不存在")
		return
	}

	if admin.GetType(ctx) > admUser.Type {
		// 不能改权限比自己高的(Type越小权限越高)
		err = errors.New("权限不足")
		return
	}

	sys.RepoAdmUser.Update(&model.AdmUser{Id: admUser.Id, Password: util.Md5(vmAdmUser.Password)})

	return
}

func (*Sys) MasterList(ctx *gin.Context) (*page.Info, []*dto.AdmUserDto) {
	var admUsers = make([]model.AdmUser, 0)
	var wheres = make([][]interface{}, 0)

	typeMe := admin.GetType(ctx)
	wheres = append(wheres, []interface{}{"type", ">=", typeMe})

	pageInfo := page.PageWhereOrder(ctx, &admUsers, wheres, "`id`")

	var admUserDtos = make([]*dto.AdmUserDto, 0)
	for _, admUser := range admUsers {
		admUserDtos = append(admUserDtos, dto.ToAdmUser(&admUser))
	}

	return pageInfo, admUserDtos
}

func (*Sys) LogLogin(ctx *gin.Context, loginVm *vm.LogAdmUserLoginVm) (*page.Info, []*dto.LogAdmUserLoginDto) {
	var userLogins = make([]model.LogAdmUserLogin, 0)
	var wheres = make([][]interface{}, 0)

	if loginVm.Account != "" {
		wheres = append(wheres, []interface{}{"account", loginVm.Account})
	}

	pageInfo := page.PageWhereOrder(ctx, &userLogins, wheres, "`id` DESC")

	var logAdmUserLoginDtos = make([]*dto.LogAdmUserLoginDto, 0)
	for _, userLogin := range userLogins {
		logAdmUserLoginDtos = append(logAdmUserLoginDtos, dto.ToLogAdmUserLogin(&userLogin))
	}

	return pageInfo, logAdmUserLoginDtos
}
