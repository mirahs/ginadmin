package main

import (
	"ginadmin/app"
	"ginadmin/app/config"
)


func main() {
	app.Start(&config.App{
		MysqlHost:     "127.0.0.1",
		MysqlPort:     3306,
		MysqlDatabase: "ginadmin",
		MysqlUser:     "root",
		MysqlPassword: "root",

		DefaultAccount:     "admin",
		DefaultPassword:    "admin",
		DefaultType:        config.AdminUserTypeAdmin,
		DefaultNewPassword: "123456",

		TemplateDir: "./app/template/",
		StaticDir:   "./app/static/",

		SessionName: "ginadmin_session",
		SessionSecret: "ginadmin_secret",

		Ip2RegionDbFile: "./app/ip2region.db",
	})
}
