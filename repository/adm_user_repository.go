package repository

import (
	"ginadmin/model"
	"ginadmin/util"
)

type AdmUser struct {}


func (*AdmUser) GetByAccount(account string) *model.AdmUser {
	var admUser model.AdmUser
	model.Db.Find(&admUser, "`account`=?", account)
	return &admUser
}

func (*AdmUser) LoginUpdate(admUser *model.AdmUser, ip string)  {
	admUser.LoginTimes++
	admUser.LoginTime = uint32(util.Unixtime())
	admUser.LoginIp = ip

	model.Db.Save(admUser)
}

func (*AdmUser) UpdatePassword(account, password string) {
	var admUser model.AdmUser
	model.Db.Model(admUser).Where("`account`=?", account).Update("password", password)
}


func NewAdmUserRepository() *AdmUser {
	return &AdmUser{}
}
