package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


func GinSuccess(ctx *gin.Context) {
	GinJson(ctx, gin.H{"code": 1})
}

func GinSuccessByData(ctx *gin.Context, obj interface{}) {
	GinJson(ctx, gin.H{"code": 1, "data": obj})
}

func GinFailed(ctx *gin.Context, msg string) {
	GinJson(ctx, gin.H{"code": 0, "msg": msg})
}

func GinJson(ctx *gin.Context, obj interface{}) {
	ctx.JSON(http.StatusOK, obj)
}

func GinError(ctx *gin.Context, msg string) {
	content := fmt.Sprintf("<script type=\"text/javascript\">alert('%s');</script>", msg)

	ctx.Header("content-type", "text/html; charset=utf-8")
	ctx.Header("refresh", fmt.Sprintf("0;url=\"%s\"", ctx.FullPath()))

	ctx.String(http.StatusOK, content)
}

func GinRedirect(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, ctx.FullPath())
}
