package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func Index(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func IndexLogin(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK, "login.html", nil)
}
