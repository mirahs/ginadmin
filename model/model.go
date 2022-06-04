package model

import (
	"fmt"
	"ginadmin/common"
	"ginadmin/conf"
	"ginadmin/util"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)


var Db *gorm.DB


// Init gorm 初始化
func Init() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.App.MysqlUser, conf.App.MysqlPassword, conf.App.MysqlHost,
		conf.App.MysqlPort, conf.App.MysqlDatabase,
	)

	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("db.Init Open err:" + err.Error())
	}

	sqlDB, err := Db.DB()
	if err != nil {
		panic("model.Init DB err:" + err.Error())
	}
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 判断 adm_user 表是否存在(要在 AutoMigrate 操作之前)
	hasAdmUser := Db.Migrator().HasTable(&AdmUser{})

	// model初始化
	err = Db.AutoMigrate(&AdmUser{}, &LogAdmUserLogin{})
	if err != nil {
		panic("db.Init AutoMigrate err:" + err.Error())
	}

	// todo: 如果启动时 adm_user 表不存在就创建初始化管理员账号(登录后记得改密码或者删除这个账号)
	if !hasAdmUser {
		Db.Create(&AdmUser{
			Account:  conf.App.InitAccount,
			Password: util.Md5(conf.App.InitPassword),
			Type:     common.AdminUserTypeAdmin,
			Remark:   conf.App.InitAccount,
		})
	}

	// 开发模式开启 gorm 调试
	if conf.App.GinMode == gin.DebugMode {
		Db = Db.Debug()
	}
}

func Close() {
	sqlDB, _ := Db.DB()
	_ = sqlDB.Close()
}
