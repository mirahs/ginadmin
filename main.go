package main

import (
	"ginadmin/app"
	"ginadmin/app/config"
)


func main() {
	// 启动app, app有默认参数, 如果跟自己的不一致, 在这里替换
	app.Start(&config.App{
		MysqlHost:     "127.0.0.1",
		MysqlPort:     3306,
		MysqlDatabase: "ginadmin",
		MysqlUser:     "root",
		MysqlPassword: "root",
	})
}
