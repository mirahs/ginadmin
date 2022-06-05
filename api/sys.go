package api

import (
	"ginadmin/common"
	"ginadmin/conf"
	"ginadmin/dto"
	"ginadmin/model"
	"ginadmin/service"
	"ginadmin/util"
	"ginadmin/vm"
	"ginadmin/vo"
	"github.com/flosch/pongo2/v4"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"sort"
)


// SysPassword 修改密码
func SysPassword(ctx *gin.Context)  {
	if ctx.Request.Method == "POST" {
		errMsg := service.SysPassword(ctx)
		if errMsg != "" {
			ctx.JSON(http.StatusOK, gin.H{"code": 1, "msg": errMsg})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"code": 0})
	} else {
		util.HTML(ctx, "sys/password.html", pongo2.Context{
			"account": service.GetAccount(ctx),
		})
	}
}

// SysMasterNew 添加管理员
func SysMasterNew(ctx *gin.Context) {
	if ctx.Request.Method == "GET" {
		context := pongo2.Context{}
		vmUser := vm.AdminUser{}

		id := uint32(com.StrTo(ctx.Query("id")).MustInt())
		if id != 0 {
			vmUser.Id = id

			if admUser, err := vmUser.Get(); err == nil {
				context["data"] = admUser
			}
		}

		userType := util.GetType(ctx)
		var userTypeDescs []*dto.UserTypeDescDto
		for key, val := range common.AdminUserTypesDesc {
			if key >= userType {
				userTypeDescs = append(userTypeDescs, dto.ToUserType(key, val))
			}
		}
		// common.AdminUserTypesDesc 是一个map, 转成切片的时候是无序的, 所以需要排序(用户类型从小到大)
		sort.Sort(dto.UTDSlice(userTypeDescs))

		context["user_types"] = userTypeDescs

		util.HTML(ctx, "sys/master_new.html", context)
	} else {
		voNew := vo.AdmUserNew{}
		err := util.GinBindValid(ctx, &voNew)
		if err != nil {
			util.GinError(ctx, "请输入正确的数据")
			return
		}

		vmUser := vm.AdminUser{Id: voNew.Id, Account: voNew.Account, Type: voNew.Type, Remark: voNew.Remark}

		if voNew.Id > 0 {
			err = vmUser.Edit()
		} else {
			vmUser.Password = conf.App.DefaultPassword
			err = vmUser.Add()
		}
		if err != nil {
			util.GinError(ctx, "DB错误")
			return
		}

		util.GinSuccess(ctx)
	}
}

// SysMasterList 管理员列表
func SysMasterList(ctx *gin.Context) {
	voDR := vo.AdmUserDelOrReset{}
	_ = ctx.ShouldBind(&voDR)

	vmUser := vm.AdminUser{}

	switch ctx.Query("act") {
	case "del":
		if voDR.Id == 0 || voDR.Id == util.GetId(ctx) {
			util.GinError(ctx,"不能删除自己")
		} else {
			vmUser.Id = voDR.Id
			_ = vmUser.Delete()
			util.GinError(ctx,"删除成功")
		}
	case "lock":
		if voDR.Id == 0 || voDR.Id == util.GetId(ctx) {
			util.GinError(ctx,"不能操作自己")
		} else {
			isLocked := uint8(util.If(voDR.IsLocked == 0, 1, 0).(int))

			vmUser.Id = voDR.Id
			vmUser.IsLocked = isLocked

			_ = vmUser.Lock()

			util.GinRedirect(ctx)
		}
	default:
		pageInfo, admUsers := service.MasterList(ctx)

		util.HTML(ctx, "sys/master_list.html", pongo2.Context{
			"page": pageInfo,
			"datas": admUsers,
		})
	}
}

// SysLogLogin 用户登录日志
func SysLogLogin(ctx *gin.Context)  {
	id := uint32(com.StrTo(ctx.Query("id")).MustInt())
	if id != 0 {
		model.LogAdmUserDelete(id)

		util.GinRedirect(ctx)
		return
	}

	var voLog vo.LogAdmUser
	_ = ctx.ShouldBind(&voLog)

	pageInfo, logLogins := service.LogLogin(ctx, &voLog)

	util.HTML(ctx, "sys/log_login.html", pongo2.Context{
		"account": voLog.Account,
		"page": pageInfo,
		"datas": logLogins,
	})
}
