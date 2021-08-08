package common

import (
	"ginadmin/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var Db *gorm.DB


func DbInit() {
	dsn := "root:root@tcp(127.0.0.1:3306)/ginadmin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("db.DbInit Open err:" + err.Error())
	}
	Db = db

	hasAdmUser := Db.Migrator().HasTable(&model.AdmUser{})

	err = Db.AutoMigrate(&model.AdmUser{}, &model.LogAdmUserLogin{})
	if err != nil {
		panic("db.DbInit AutoMigrate err:" + err.Error())
	}

	if !hasAdmUser {
		Db.Create(&model.AdmUser{Account: "admin", Password: "21232f297a57a5a743894a0e4a801fc3", Type: 10, Remark: "admin"})
	}
}
