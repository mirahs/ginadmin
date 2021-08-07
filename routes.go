package main

import (
	"ginadmin/api/admin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)


func initRoutes(engine *gin.Engine) {
	adminG := engine.Group("/admin")

	initSession(adminG)

	adminG.GET("/index/index", admin.Index)
	adminG.GET("/index/login", admin.IndexLogin)
	adminG.POST("/index/login", admin.IndexLogin)
	adminG.GET("/index/logout", admin.IndexLogout)
}

func initSession(group *gin.RouterGroup) {
	store := memstore.NewStore([]byte("ginadminsecret"))
	group.Use(sessions.Sessions("ginadminsession", store))
}
