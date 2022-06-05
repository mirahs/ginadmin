package service

import (
	"ginadmin/common"
	"ginadmin/model"
	"ginadmin/util"
	"ginadmin/vo"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)


// 登录成功后 session 设置
func loginSessionSet(ctx *gin.Context, admUser *model.AdmUser) {
	session := sessions.Default(ctx)

	session.Set(common.SessionId, admUser.Id)
	session.Set(common.SessionAccount, admUser.Account)
	session.Set(common.SessionType, admUser.Type)
	session.Set(common.SessionIsLogin, true)

	_ = session.Save()
}

func GetAccount(ctx *gin.Context) string {
	var voAdmUser vo.AdmUserAccount
	_ = ctx.ShouldBind(&voAdmUser)

	if voAdmUser.Account != "" {
		return voAdmUser.Account
	}

	return util.GetAccount(ctx)
}
