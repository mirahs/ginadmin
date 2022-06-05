package util

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)


var validate = validator.New()

// GinBindValid 绑定并验证
func GinBindValid(ctx *gin.Context, bind interface{}) error {
	err := ctx.ShouldBind(bind)
	if err != nil {
		return err
	}

	err = validate.Struct(bind)

	if err != nil {
		errFirst := err.(validator.ValidationErrors)[0]
		return errors.New(fmt.Sprintf("%s Value:%v", errFirst, errFirst.Value()))
	}

	return nil
}


// json 返回成功
func GinSuccess(ctx *gin.Context) {
	GinJson(ctx, gin.H{"code": 1})
}

// json 返回成功(有数据)
func GinSuccessByData(ctx *gin.Context, obj interface{}) {
	GinJson(ctx, gin.H{"code": 1, "data": obj})
}

// json 返回失败
func GinFailed(ctx *gin.Context, msg string) {
	GinJson(ctx, gin.H{"code": 0, "msg": msg})
}

// json 返回 code、msg
func GinJson(ctx *gin.Context, obj interface{}) {
	ctx.JSON(http.StatusOK, obj)
}

// alert 错误
func GinError(ctx *gin.Context, msg string) {
	content := fmt.Sprintf("<script type=\"text/javascript\">alert('%s');</script>", msg)

	ctx.Header("content-type", "text/html; charset=utf-8")
	ctx.Header("refresh", fmt.Sprintf("0;url=\"%s\"", ctx.FullPath()))

	ctx.String(http.StatusOK, content)
}

// 默认跳转到当前页面
func GinRedirect(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, ctx.FullPath())
}
