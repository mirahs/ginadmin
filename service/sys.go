package service

import (
	"ginadmin/dto"
	"ginadmin/model"
	"ginadmin/util"
	"ginadmin/util/page"
	"ginadmin/vm"
	"ginadmin/vo"
	"github.com/gin-gonic/gin"
)


// SysPassword 更改密码
func SysPassword(ctx *gin.Context) string {
	var voAdmUser vo.AdmUserPassword

	err := util.GinBindValid(ctx, &voAdmUser)
	if err != nil {
		return "参数错误"
	}

	vmUser := vm.AdminUser{Account: voAdmUser.Account, Password: voAdmUser.Password}

	admUser, err := vmUser.GetByAccount()
	if err != nil {
		return "DB错误"
	}
	if admUser.Id == 0 {
		return "账号不存在"
	}

	if util.GetType(ctx) > admUser.Type {
		// 不能改权限比自己高的(Type越小权限越高)
		return "权限不足"
	}

	vmUser.Id = admUser.Id

	err = vmUser.PasswordFun()
	if err != nil {
		return "DB错误"
	}

	return ""
}

func MasterList(ctx *gin.Context) (*page.Info, []*dto.AdmUserDto) {
	var admUsers []*model.AdmUser
	var wheres [][]interface{}
	orders := []string{"`id`"}

	typeMe := util.GetType(ctx)
	wheres = append(wheres, []interface{}{"`type`>=?", typeMe})

	pageInfo, _ := page.Page(&page.Param{Ctx: ctx, Db: model.Db, Datas: &admUsers, Wheres: wheres, Orders: orders})

	var admUserDtos = make([]*dto.AdmUserDto, 0)
	for _, admUser := range admUsers {
		admUserDtos = append(admUserDtos, dto.ToAdmUser(admUser))
	}

	return pageInfo, admUserDtos
}

func LogLogin(ctx *gin.Context, loginVo *vo.LogAdmUser) (*page.Info, []*dto.LogAdmUserLoginDto) {
	var userLogins []*model.LogAdmUserLogin
	var wheres [][]interface{}
	preloads := [][]interface{}{{"AdmUser"}}
	orders := []string{"`id` DESC"}

	if loginVo.Account != "" {
		vmUser := vm.AdminUser{Account: loginVo.Account}
		if user, err := vmUser.GetByAccount(); err == nil && user.Id != 0 {
			wheres = append(wheres, []interface{}{"`adm_user_id`=?", user.Id})
		} else {
			wheres = append(wheres, []interface{}{"`adm_user_id`=0"})
		}
	}

	pageInfo, _ := page.Page(&page.Param{Ctx: ctx, Db: model.Db, Datas: &userLogins, Wheres: wheres, Preloads: preloads, Orders: orders})

	var logAdmUserLoginDtos []*dto.LogAdmUserLoginDto
	for _, userLogin := range userLogins {
		logAdmUserLoginDtos = append(logAdmUserLoginDtos, dto.ToLogAdmUserLogin(userLogin))
	}

	return pageInfo, logAdmUserLoginDtos
}
