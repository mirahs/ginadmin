package model

import (
	"fmt"
	"ginadmin/app/config"
	"ginadmin/app/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var Db *gorm.DB


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

	hasAdmUser := Db.Migrator().HasTable(&AdmUser{})

	err = Db.AutoMigrate(&AdmUser{}, &LogAdmUserLogin{})
	if err != nil {
		panic("db.DbInit AutoMigrate err:" + err.Error())
	}

	if !hasAdmUser {
		Db.Create(&AdmUser{
			Account: config.AppInst.InitAccount,
			Password: util.Md5(config.AppInst.InitPassword),
			Type: config.AdminUserTypeAdmin,
			Remark: config.AppInst.InitAccount,
		})
	}
}
