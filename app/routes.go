package app

import (
	"ginadmin/app/api/admin"
	"ginadmin/app/middleware"
	"github.com/gin-gonic/gin"
)


// 初始化路由
func initRoutes(engine *gin.Engine) {
	adminG := engine.Group("/admin")
	{
		initSession(adminG)

		adminG.Use(middleware.AdminValidate())

		adminG.GET("/index/index", admin.Index)
		adminG.GET("/index/login", admin.IndexLogin)
		adminG.POST("/index/login", admin.IndexLogin)
		adminG.GET("/index/logout", admin.IndexLogout)
		adminG.GET("/index/deny", admin.IndexDeny)

		adminG.GET("/home/welcome", admin.HomeWelcome)

		adminG.GET("/sys/password", admin.SysPassword)
		adminG.POST("/sys/password", admin.SysPassword)
		adminG.GET("/sys/master_new", admin.SysMasterNew)
		adminG.POST("/sys/master_new", admin.SysMasterNew)
		adminG.GET("/sys/master_list", admin.SysMasterList)
		adminG.GET("/sys/log_login", admin.SysLogLogin)
	}
}
