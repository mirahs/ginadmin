package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func HomeWelcome(ctx *gin.Context) {
	ctx.String(http.StatusOK, "welcome")
}
