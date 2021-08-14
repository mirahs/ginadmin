package admin

import (
	"ginadmin/config"
	"ginadmin/dto"
	"ginadmin/repository"
	admin2 "ginadmin/service/admin"
	"ginadmin/util"
	"ginadmin/util/admin"
	"github.com/flosch/pongo2/v4"
	"github.com/gin-gonic/gin"
	"net/http"
)


var sysService = admin2.NewSysService()
var repoAdmUser = repository.NewAdmUserRepository()


func SysPassword(ctx *gin.Context)  {
	if ctx.Request.Method == "POST" {
		err := sysService.Password(ctx)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"code": 1})
	} else {
		vmAdmUser := sysService.AccountInfo(ctx)

		ctx.HTML(http.StatusOK, "admin/sys/password.html", pongo2.Context{
			"account": vmAdmUser.Account,
		})
	}
}

func SysMasterNew(ctx *gin.Context) {
	vmAdmUser := sysService.BindAdmUser(ctx)
	context := pongo2.Context{}

	if ctx.Request.Method == "GET" {
		if vmAdmUser.Id != 0 {
			admUser := repoAdmUser.Get(vmAdmUser.Id)
			context["data"] = admUser
		}

		userType := admin.GetAccountType(ctx)
		var userTypeDescs []*dto.UserTypeDescDto
		for key, val := range config.AdminUserTypesDesc {
			if key >= userType {
				userTypeDescs = append(userTypeDescs, dto.ToUserType(key, val))
			}
		}
		context["user_types"] = userTypeDescs

		ctx.HTML(http.StatusOK, "admin/sys/master_new.html", context)
	} else {
		if vmAdmUser.Account == "" || vmAdmUser.Type == 0 || vmAdmUser.Remark == "" {
			util.GinError(ctx, "请输入正确的数据")
			return
		}

		admUser := sysService.AdmUserVm2AdmUser(vmAdmUser)
		if admUser.Id > 0 {
			repoAdmUser.Update(admUser)
		} else {
			admUser.Password = util.Md5("123456")
			repoAdmUser.Add(admUser)
		}
		util.GinSuccess(ctx)
	}
}

func SysMasterList(ctx *gin.Context) {
	vmAdmUser := sysService.BindAdmUser(ctx)

	switch ctx.Query("act") {
	case "del":
		if vmAdmUser.Id == admin.GetAccountId(ctx) {
			util.GinError(ctx,"不能删除自己")
		} else {
			repoAdmUser.DelById(vmAdmUser.Id)
			util.GinError(ctx,"删除成功")
		}
	case "lock":
		if vmAdmUser.Id == admin.GetAccountId(ctx) {
			util.GinError(ctx,"不能操作自己")
		} else {
			isLocked := uint8(util.If(vmAdmUser.IsLocked == 0, 1, 0).(int))
			repoAdmUser.UpdateIsLockedById(vmAdmUser.Id, isLocked)
			util.GinRedirect(ctx)
		}
	default:
		pageInfo, admUsers := sysService.MasterList(ctx)

		ctx.HTML(http.StatusOK, "admin/sys/master_list.html", pongo2.Context{
			"page": pageInfo,
			"datas": admUsers,
		})
	}
}
