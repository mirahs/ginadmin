package main

import (
	"ginadmin/api/admin"
	"github.com/gin-gonic/gin"
)


func registerRoutes(engine *gin.Engine)  {
	adminG := engine.Group("/admin")

	adminG.GET("/index", admin.Index)
	adminG.GET("/index/login", admin.IndexLogin)
	adminG.POST("/index/login", admin.IndexLogin)
}
