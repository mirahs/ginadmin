package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


// 欢迎页
func HomeWelcome(ctx *gin.Context) {
	ctx.String(http.StatusOK, "welcome")
}
