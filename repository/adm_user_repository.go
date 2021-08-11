package repository

import (
	"ginadmin/model"
)

type AdmUser struct {}


func (*AdmUser) GetByAccount(account string) *model.AdmUser {
	var admUser model.AdmUser
	model.Db.Find(&admUser, "`account`=?", account)
	return &admUser
}


func NewAdmUserRepository() *AdmUser {
	return &AdmUser{}
}
