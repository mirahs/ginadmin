package admin

import (
	"ginadmin/common"
	"ginadmin/conf"
	"ginadmin/dto"
	"ginadmin/util"
	"ginadmin/util/admin"
	"ginadmin/vo"
	"github.com/flosch/pongo2/v4"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
)


// 修改密码
func SysPassword(ctx *gin.Context)  {
	if ctx.Request.Method == "POST" {
		err := serviceSys.Password(ctx)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"code": 1})
	} else {
		voAdmUser := serviceSys.AccountInfo(ctx)

		admin.HTML(ctx, "admin/sys/password.html", pongo2.Context{
			"account": voAdmUser.Account,
		})
	}
}

// 添加管理员
func SysMasterNew(ctx *gin.Context) {
	voAdmUser := serviceSys.BindAdmUser(ctx)
	context := pongo2.Context{}

	if ctx.Request.Method == "GET" {
		if voAdmUser.Id != 0 {
			admUser := serviceSys.RepoAdmUser.Get(voAdmUser.Id)
			context["data"] = admUser
		}

		userType := admin.GetType(ctx)
		var userTypeDescs []*dto.UserTypeDescDto
		for key, val := range common.AdminUserTypesDesc {
			if key >= userType {
				userTypeDescs = append(userTypeDescs, dto.ToUserType(key, val))
			}
		}
		// common.AdminUserTypesDesc 是一个map, 转成切片的时候是无序的, 所以需要排序(用户类型从小到大)
		sort.Sort(dto.UTDSlice(userTypeDescs))

		context["user_types"] = userTypeDescs

		admin.HTML(ctx, "admin/sys/master_new.html", context)
	} else {
		if voAdmUser.Account == "" || voAdmUser.Type == 0 || voAdmUser.Remark == "" {
			util.GinError(ctx, "请输入正确的数据")
			return
		}

		admUser := serviceSys.AdmUserVo2AdmUser(voAdmUser)
		if admUser.Id > 0 {
			serviceSys.RepoAdmUser.Update(admUser)
		} else {
			admUser.Password = util.Md5(conf.App.DefaultPassword)
			serviceSys.RepoAdmUser.Add(admUser)
		}
		util.GinSuccess(ctx)
	}
}

// 管理员列表
func SysMasterList(ctx *gin.Context) {
	voAdmUser := serviceSys.BindAdmUser(ctx)

	switch ctx.Query("act") {
	case "del":
		if voAdmUser.Id == admin.GetId(ctx) {
			util.GinError(ctx,"不能删除自己")
		} else {
			serviceSys.RepoAdmUser.DelById(voAdmUser.Id)
			util.GinError(ctx,"删除成功")
		}
	case "lock":
		if voAdmUser.Id == admin.GetId(ctx) {
			util.GinError(ctx,"不能操作自己")
		} else {
			isLocked := uint8(util.If(voAdmUser.IsLocked == 0, 1, 0).(int))
			serviceSys.RepoAdmUser.UpdateIsLockedById(voAdmUser.Id, isLocked)
			util.GinRedirect(ctx)
		}
	default:
		pageInfo, admUsers := serviceSys.MasterList(ctx)

		admin.HTML(ctx, "admin/sys/master_list.html", pongo2.Context{
			"page": pageInfo,
			"datas": admUsers,
		})
	}
}

// 用户登录日志
func SysLogLogin(ctx *gin.Context)  {
	var voUserLogin vo.LogAdmUserLoginVo
	_ = ctx.ShouldBind(&voUserLogin)

	if voUserLogin.Id != 0 {
		serviceSys.RepoLogAdmUserLogin.DelById(voUserLogin.Id)
		util.GinRedirect(ctx)
		return
	}

	pageInfo, logLogins := serviceSys.LogLogin(ctx, &voUserLogin)

	admin.HTML(ctx, "admin/sys/log_login.html", pongo2.Context{
		"account": voUserLogin.Account,
		"page": pageInfo,
		"datas": logLogins,
	})
}
