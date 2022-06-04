package admin

import (
	"errors"
	"ginadmin/util"
	"ginadmin/util/admin"
	"ginadmin/vo"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)


type Index struct {
	base
}


func (index *Index) Login(ctx *gin.Context) (err error)  {
	var voAdmUser vo.AdmUserVo
	err = ctx.ShouldBind(&voAdmUser)
	if err != nil {
		err = errors.New("参数解析错误:" + err.Error())
		return
	}

	if voAdmUser.Account == "" || voAdmUser.Password == "" {
		err = errors.New("账号和密码不能为空")
		return
	}

	ip := ctx.ClientIP()
	address := util.IpAddress(ip)

	var remark string

	admUser := index.RepoAdmUser.GetByAccount(voAdmUser.Account)
	if admUser.Id == 0 {
		remark = "账号不存在"
		err = errors.New(remark)
		index.RepoLogAdmUserLogin.AddFailed(admUser.Id, ip, address, remark)
		return
	}

	passwordMd5 := util.Md5(voAdmUser.Password)
	if passwordMd5 != admUser.Password {
		remark = "密码错误"
		err = errors.New(remark)
		index.RepoLogAdmUserLogin.AddFailed(admUser.Id, ip, address, remark)
		return
	}

	admin.LoginSessionSet(ctx, admUser)

	index.RepoAdmUser.LoginUpdate(admUser, ip)
	index.RepoLogAdmUserLogin.AddSuccess(admUser.Id, ip, address, remark)

	return
}

func (*Index) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	_ = session.Save()
}
