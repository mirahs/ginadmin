package repository

import (
	"ginadmin/model"
	"ginadmin/util"
)


type LogAdmUserLogin struct{}


func (login *LogAdmUserLogin) AddSuccess(accountId uint32, ip string, address string, remark string) {
	login.add(accountId, 1, ip, address, remark)
}

func (login *LogAdmUserLogin) AddFailed(accountId uint32, ip string, address string, remark string) {
	login.add(accountId, 0, ip, address, remark)
}

func (*LogAdmUserLogin) DelById(id uint32) {
	var logAdmUserLogin model.LogAdmUserLogin
	model.Db.Where("`id`=?", id).Delete(logAdmUserLogin)
}

func (*LogAdmUserLogin) add(accountId uint32, status int, ip string, address string, remark string) {
	model.Db.Create(&model.LogAdmUserLogin{
		AdmUserID: accountId,
		Time:    uint32(util.Unixtime()),
		Status:  uint8(status),
		Ip:      ip,
		Address: address,
		Remark:  remark,
	})
}
