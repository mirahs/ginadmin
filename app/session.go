package app

import (
	"ginadmin/app/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)


func initSession(group *gin.RouterGroup) {
	store := memstore.NewStore([]byte(config.AppInst.SessionSecret))
	store.Options(sessions.Options{Path: group.BasePath()})

	group.Use(sessions.Sessions(config.AppInst.SessionName, store))
}
