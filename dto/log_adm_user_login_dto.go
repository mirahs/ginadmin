package dto

import (
	"ginadmin/model"
	"ginadmin/util"
)


type LogAdmUserLoginDto struct {
	Id       uint32
	Account  string
	TimeDesc string
	StatusDesc string
	Ip string
	Address string
	Remark string
}


func ToLogAdmUserLogin(login *model.LogAdmUserLogin) *LogAdmUserLoginDto {
	return &LogAdmUserLoginDto{
		Id:         login.Id,
		Account:    login.AdmUser.Account,
		TimeDesc:   util.Time2Datetime(int64(login.Time)),
		StatusDesc: util.If(login.Status == 0, "失败", "成功").(string),
		Ip:         login.Ip,
		Address:    login.Address,
		Remark:     login.Remark,
	}
}
