package admin

import (
	"ginadmin/app/config"
	"ginadmin/app/model"
	"github.com/flosch/pongo2/v4"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)


// 获取用户 ID
func GetId(ctx *gin.Context) uint32 {
	session := sessions.Default(ctx)
	return session.Get(config.SessionId).(uint32)
}

// 获取用户账号
func GetAccount(ctx *gin.Context) string {
	session := sessions.Default(ctx)
	return session.Get(config.SessionAccount).(string)
}

// 获取用户类型
func GetType(ctx *gin.Context) uint8 {
	session := sessions.Default(ctx)
	return session.Get(config.SessionType).(uint8)
}


// 登录成功后 session 设置
func LoginSessionSet(ctx *gin.Context, admUser *model.AdmUser) {
	session := sessions.Default(ctx)

	session.Set(config.SessionId, admUser.Id)
	session.Set(config.SessionAccount, admUser.Account)
	session.Set(config.SessionType, admUser.Type)
	session.Set(config.SessionIsLogin, true)

	_ = session.Save()
}

// 登录检查
func LoginCheck(ctx *gin.Context) bool {
	session := sessions.Default(ctx)
	isLogin := session.Get(config.SessionIsLogin)
	return isLogin != nil && isLogin.(bool)
}


// 模板渲染
func HTML(ctx *gin.Context, tplPath string, obj pongo2.Context)  {
	if obj == nil {
		obj = pongo2.Context{}
	}

	obj["static_url"] = config.AppInst.StaticUrl

	ctx.HTML(http.StatusOK, tplPath, obj)
}
