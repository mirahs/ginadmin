package admin

import (
	"github.com/flosch/pongo2/v4"
	"github.com/gin-gonic/gin"
	"net/http"
)


func Index(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK, "admin/index/index.html", pongo2.Context{})
}

func IndexLogin(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK, "admin/index/login.html", pongo2.Context{})
}
