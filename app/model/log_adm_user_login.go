package model


type LogAdmUserLogin struct {
	Id      uint32 `gorm:"auto_increment; comment:Id"`
	Account string `gorm:"type:varchar(32); not null; comment:帐号"`
	Time    uint32 `gorm:"not null; comment:时间"`
	Status  uint8  `gorm:"not null; comment:状态 0:失败|1:成功"`
	Ip      string `gorm:"type:varchar(15); not null; comment:IP"`
	Address string `gorm:"type:varchar(128); not null; comment:地址"`
	Remark  string `gorm:"type:varchar(32); not null; comment:备注"`
}

// 设置表名，gorm 默认是复数形式
func (user *LogAdmUserLogin) TableName() string {
	return "log_adm_user_login"
}
