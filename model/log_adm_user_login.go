package model

import "ginadmin/util"

type LogAdmUserLogin struct {
	Id      uint32

	AdmUserID uint32 `gorm:"not null; comment:管理员ID; index"`
	AdmUser *AdmUser
	Time    uint32 `gorm:"not null; comment:时间; index"`
	Status  uint8  `gorm:"not null; comment:状态 0:失败|1:成功"`
	Ip      string `gorm:"type:varchar(15); not null; comment:IP"`
	Address string `gorm:"type:varchar(128); not null; comment:地址"`
	Remark  string `gorm:"type:varchar(32); not null; comment:备注"`
}

// TableName 设置表名，gorm 默认是复数形式
func (user *LogAdmUserLogin) TableName() string {
	return "log_adm_user_login"
}


func LogAdmUserLoginAddSuccess(admUserId uint32, ip, address, remark string) {
	logAdmUserLoginAdd(admUserId, ip, address, remark, 1)
}

func LogAdmUserLoginAddFailed(admUserId uint32, ip, address, remark string) {
	logAdmUserLoginAdd(admUserId, ip, address, remark, 0)
}

func LogAdmUserDelete(id uint32) {
	Db.Delete(&LogAdmUserLogin{Id: id})
}


func logAdmUserLoginAdd(admUserId uint32, ip, address, remark string, status int) {
	Db.Create(&LogAdmUserLogin{
		AdmUserID: admUserId,
		Time:    uint32(util.UnixTime()),
		Status:  uint8(status),
		Ip:      ip,
		Address: address,
		Remark:  remark,
	})
}
