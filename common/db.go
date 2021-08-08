package common

import (
	"ginadmin/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var Db *gorm.DB


func DbInit()  {
	dsn := "root:root@tcp(127.0.0.1:3306)/ginadmin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("db.DbInit Open err:" + err.Error())
	}
	Db = db

	db.AutoMigrate(&model.AdmUser{})
	db.AutoMigrate(&model.LogAdmUserLogin{})
}
