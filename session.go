package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)


func initSession(group *gin.RouterGroup) {
	store := memstore.NewStore([]byte("ginadminsecret"))
	group.Use(sessions.Sessions("ginadminsession", store))
}
