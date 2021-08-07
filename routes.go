package main

import (
	"ginadmin/api/admin"
	"github.com/gin-gonic/gin"
)


func registerRoutes(engine *gin.Engine)  {
	adm := engine.Group("/admin")
	adm.GET("/index", admin.Index)
	adm.Any("/index/login", admin.IndexLogin)
}
