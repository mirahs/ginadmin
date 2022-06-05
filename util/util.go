package util

import (
	"crypto/md5"
	"fmt"
	"ginadmin/common"
	"ginadmin/conf"
	"github.com/flosch/pongo2/v4"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"time"
)


// 1970年01月01日00时00分00秒起至现在的秒数
func UnixTime() int64 {
	return time.Now().Unix()
}

// 时间戳转年月日时分秒
func Time2Datetime(unixtime int64) string {
	return time.Unix(unixtime, 0).Format("2006-01-02 15:04:05")
}

// md5 散列
func Md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

// 三目运算
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}


func PageInfo(ctx *gin.Context) (int, int) {
	page := com.StrTo(ctx.Query("page")).MustInt()
	pageSize := com.StrTo(ctx.Query("limit")).MustInt()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	return page, pageSize
}


// 获取用户 ID
func GetId(ctx *gin.Context) uint32 {
	session := sessions.Default(ctx)
	return session.Get(common.SessionId).(uint32)
}

// 获取用户账号
func GetAccount(ctx *gin.Context) string {
	session := sessions.Default(ctx)
	return session.Get(common.SessionAccount).(string)
}

// 获取用户类型
func GetType(ctx *gin.Context) uint8 {
	session := sessions.Default(ctx)
	return session.Get(common.SessionType).(uint8)
}


// 登录检查
func LoginCheck(ctx *gin.Context) bool {
	session := sessions.Default(ctx)
	isLogin := session.Get(common.SessionIsLogin)
	return isLogin != nil && isLogin.(bool)
}


// 模板渲染
func HTML(ctx *gin.Context, tplPath string, obj pongo2.Context)  {
	if obj == nil {
		obj = pongo2.Context{}
	}

	obj["static_url"] = conf.App.StaticUrl

	ctx.HTML(http.StatusOK, tplPath, obj)
}

