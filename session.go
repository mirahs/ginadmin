package main

import (
	"ginadmin/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)


// 初始化 session
func initSession(group *gin.RouterGroup) {
	store := memstore.NewStore([]byte(config.App.SessionSecret))
	store.Options(sessions.Options{Path: group.BasePath()})

	group.Use(sessions.Sessions(config.App.SessionName, store))
}
