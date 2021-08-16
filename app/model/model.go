package model

import (
	"fmt"
	"ginadmin/app/config"
	"ginadmin/app/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var Db *gorm.DB


// gorm初始化
func DbInit() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.AppInst.MysqlUser, config.AppInst.MysqlPassword, config.AppInst.MysqlHost,
		config.AppInst.MysqlPort, config.AppInst.MysqlDatabase,
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
			Account: config.AppInst.InitAccount,
			Password: util.Md5(config.AppInst.InitPassword),
			Type: config.AdminUserTypeAdmin,
			Remark: config.AppInst.InitAccount,
		})
	}
}
