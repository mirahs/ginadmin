package model

import (
	"fmt"
	"ginadmin/common"
	"ginadmin/config"
	"ginadmin/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var Db *gorm.DB


// gorm 初始化
func DbInit() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.App.MysqlUser, config.App.MysqlPassword, config.App.MysqlHost,
		config.App.MysqlPort, config.App.MysqlDatabase,
	)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("db.DbInit Open err:" + err.Error())
	}

	// 判断 adm_user 表是否存在(要在 AutoMigrate 操作之前)
	hasAdmUser := Db.Migrator().HasTable(&AdmUser{})

	// model初始化
	err = Db.AutoMigrate(&AdmUser{}, &LogAdmUserLogin{})
	if err != nil {
		panic("db.DbInit AutoMigrate err:" + err.Error())
	}

	// 如果启动时 adm_user 表不存在就创建初始化管理员账号(登录后记得改密码或者删除这个账号)
	if !hasAdmUser {
		Db.Create(&AdmUser{
			Account:  config.App.InitAccount,
			Password: util.Md5(config.App.InitPassword),
			Type:     common.AdminUserTypeAdmin,
			Remark:   config.App.InitAccount,
		})
	}
}
