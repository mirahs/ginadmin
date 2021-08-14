package admin

import (
	"ginadmin/config"
	"ginadmin/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)


func GetAccount(ctx *gin.Context) string {
	session := sessions.Default(ctx)
	return session.Get(config.SessionAccount).(string)
}

func GetAccountId(ctx *gin.Context) uint32 {
	session := sessions.Default(ctx)
	return session.Get(config.SessionId).(uint32)
}

func GetAccountType(ctx *gin.Context) uint8 {
	session := sessions.Default(ctx)
	return session.Get(config.SessionType).(uint8)
}


// 登录后 session 设置
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
