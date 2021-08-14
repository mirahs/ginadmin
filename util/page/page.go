package page

import (
	"fmt"
	"ginadmin/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
)


// 分页vm
type pageVm struct {
	Page int `form:"page"` //当前页
	Limit int `form:"limit"` //每页显示数量
}

// 分页信息
type Info struct {
	Curr int
	Limit int
	Count int64
	Query string
}


func Page(ctx *gin.Context, modelDatas interface{}) *Info {
	var wheres map[string]interface{}
	return work(ctx, modelDatas, wheres)
}


func work(ctx *gin.Context, modelDatas interface{}, wheres map[string]interface{}) *Info {
	page, limit := getPageLimit(ctx)

	db := model.Db
	db = formatWhere(db, wheres)

	var count int64
	db.Model(modelDatas).Count(&count)

	db.Offset((page - 1) * limit).Limit(limit).Find(modelDatas)

	return &Info{
		Curr: page,
		Limit: limit,
		Count: count,
		Query: getQuery(ctx),
	}
}

func getPageLimit(ctx *gin.Context) (int, int) {
	var pageVm pageVm

	_ = ctx.ShouldBind(&pageVm)

	if pageVm.Page <= 0 {
		pageVm.Page = 1
	}
	if pageVm.Limit <= 0 {
		pageVm.Limit = 10
	}

	return pageVm.Page, pageVm.Limit
}

func formatWhere(db *gorm.DB, wheres map[string]interface{}) *gorm.DB {
	for key, val := range wheres {
		db = db.Where(fmt.Sprintf("`%s`=?", key), val)
	}
	return db
}

func getQuery(ctx *gin.Context) string {
	var datas []string
	query := ctx.Request.URL.Query()

	for key, val := range query {
		datas = append(datas, fmt.Sprintf("%s=%s", key, val[0]))
	}

	return "&" + strings.Join(datas, "&")
}
