package main

import (
	"ginadmin/app"
	"ginadmin/app/config"
	"github.com/gin-gonic/gin"
)


func main() {
	app.Start(&config.App{
		GinMode: gin.DebugMode,
		GinPort: 8080,

		MysqlHost:     "127.0.0.1",
		MysqlPort:     3306,
		MysqlDatabase: "ginadmin",
		MysqlUser:     "root",
		MysqlPassword: "root",

		InitAccount:     "admin",
		InitPassword:    "admin",
		DefaultPassword: "123456",
	})
}
