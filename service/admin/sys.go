package admin

import (
	"errors"
	"ginadmin/dto"
	"ginadmin/model"
	"ginadmin/util"
	"ginadmin/util/admin"
	"ginadmin/util/page"
	"ginadmin/vo"
	"github.com/gin-gonic/gin"
)


type Sys struct{
	base
}


func (*Sys) BindAdmUser(ctx *gin.Context) *vo.AdmUserVo {
	var voAdmUser vo.AdmUserVo
	_ = ctx.ShouldBind(&voAdmUser)
	return &voAdmUser
}

func (sys *Sys) AccountInfo(ctx *gin.Context) *vo.AdmUserVo {
	voAdmUser := sys.BindAdmUser(ctx)

	if voAdmUser.Account == "" {
		voAdmUser.Account = admin.GetAccount(ctx)
	}

	return voAdmUser
}

func (*Sys) AdmUserVo2AdmUser(userVo *vo.AdmUserVo) *model.AdmUser {
	return &model.AdmUser{
		Id:      userVo.Id,
		Account: userVo.Account,
		Type:    userVo.Type,
		Remark:  userVo.Remark,
	}
}

// 更改密码
func (sys *Sys) Password(ctx *gin.Context) (err error) {
	var voAdmUser vo.AdmUserVo
	err = ctx.ShouldBind(&voAdmUser)
	if err != nil {
		err = errors.New("参数解析错误:" + err.Error())
		return
	}

	if voAdmUser.Account == "" || voAdmUser.Password == "" {
		err = errors.New("账号和新密码不能为空")
		return
	}

	admUser := sys.RepoAdmUser.GetByAccount(voAdmUser.Account)
	if admUser.Id == 0 {
		err = errors.New("账号不存在")
		return
	}

	if admin.GetType(ctx) > admUser.Type {
		// 不能改权限比自己高的(Type越小权限越高)
		err = errors.New("权限不足")
		return
	}

	sys.RepoAdmUser.Update(&model.AdmUser{Id: admUser.Id, Password: util.Md5(voAdmUser.Password)})

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

func (*Sys) LogLogin(ctx *gin.Context, loginVo *vo.LogAdmUserLoginVo) (*page.Info, []*dto.LogAdmUserLoginDto) {
	var userLogins = make([]model.LogAdmUserLogin, 0)
	var wheres = make([][]interface{}, 0)

	if loginVo.Account != "" {
		wheres = append(wheres, []interface{}{"account", loginVo.Account})
	}

	pageInfo := page.PageWhereOrder(ctx, &userLogins, wheres, "`id` DESC")

	var logAdmUserLoginDtos = make([]*dto.LogAdmUserLoginDto, 0)
	for _, userLogin := range userLogins {
		logAdmUserLoginDtos = append(logAdmUserLoginDtos, dto.ToLogAdmUserLogin(&userLogin))
	}

	return pageInfo, logAdmUserLoginDtos
}
