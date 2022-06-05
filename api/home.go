package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func HomeWelcome(ctx *gin.Context) {
	ctx.String(http.StatusOK, "你好")
}
