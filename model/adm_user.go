package model


type AdmUser struct {
	Id       uint32 `gorm:"auto_increment; comment:Id"`
	Account  string `gorm:"type:varchar(32); not null; comment:帐号; uniqueIndex"`
	Password string `gorm:"type:char(32); not null; comment:密码"`
	Type     uint8  `gorm:"not null; comment:类型 10:管理员|20:游客"`

	IsLocked uint8  `gorm:"not null; default:0; comment:是否被锁住 0:否|1:是"`
	Remark   string `gorm:"type:text; not null; comment:备注"`

	LoginTimes uint32 `gorm:"not null; default:0; comment:登录次数"`
	LoginTime  uint32 `gorm:"not null; default:0; comment:登录时间"`
	LoginIp    string `gorm:"type:varchar(15); not null; default:''; comment:登录IP"`
}


// 设置表名，gorm 默认是复数形式
func (user *AdmUser) TableName() string {
	return "adm_user"
}
