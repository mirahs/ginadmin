package repository

import (
	"ginadmin/app/model"
	"ginadmin/app/util"
)


type AdmUser struct {}


func (*AdmUser) Add(user *model.AdmUser) {
	model.Db.Create(user)
}

func (*AdmUser) Get(id uint32) *model.AdmUser {
	var admUser model.AdmUser
	model.Db.Find(&admUser, id)
	return &admUser
}

func (*AdmUser) GetByAccount(account string) *model.AdmUser {
	var admUser model.AdmUser
	model.Db.Find(&admUser, "`account`=?", account)
	return &admUser
}

func (*AdmUser) Update(user *model.AdmUser) {
	model.Db.Model(user).Updates(user)
}

func (*AdmUser) LoginUpdate(admUser *model.AdmUser, ip string)  {
	admUser.LoginTimes++
	admUser.LoginTime = uint32(util.Unixtime())
	admUser.LoginIp = ip

	model.Db.Save(admUser)
}

func (*AdmUser) UpdatePasswordByAccount(account, password string) {
	var admUser model.AdmUser
	model.Db.Model(admUser).Where("`account`=?", account).Update("password", password)
}

func (*AdmUser) DelById(id uint32) {
	var admUser model.AdmUser
	model.Db.Where("`id`=?", id).Delete(admUser)
}

func (*AdmUser) UpdateIsLockedById(id uint32, isLocked uint8) {
	var admUser model.AdmUser
	model.Db.Model(admUser).Where("`id`=?", id).Update("is_locked", isLocked)
}
