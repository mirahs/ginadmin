package repository

import (
	"ginadmin/app/model"
	"ginadmin/app/util"
)


type LogAdmUserLogin struct{}


func (login *LogAdmUserLogin) AddSuccess(account string, ip string, address string, remark string) {
	login.add(account, 1, ip, address, remark)
}

func (login *LogAdmUserLogin) AddFailed(account string, ip string, address string, remark string) {
	login.add(account, 0, ip, address, remark)
}

func (*LogAdmUserLogin) DelById(id uint32) {
	var logAdmUserLogin model.LogAdmUserLogin
	model.Db.Where("`id`=?", id).Delete(logAdmUserLogin)
}

func (*LogAdmUserLogin) add(account string, status int, ip string, address string, remark string) {
	model.Db.Create(&model.LogAdmUserLogin{
		Account: account,
		Time:    uint32(util.Unixtime()),
		Status:  uint8(status),
		Ip:      ip,
		Address: address,
		Remark:  remark,
	})
}


func NewLogAdmUserLoginRepository() *LogAdmUserLogin {
	return &LogAdmUserLogin{}
}
