package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var Db *gorm.DB


func DbInit() {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/ginadmin?charset=utf8mb4&parseTime=True&loc=Local"
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
		Db.Create(&AdmUser{Account: "admin", Password: "21232f297a57a5a743894a0e4a801fc3", Type: 10, Remark: "admin"})
	}
}
