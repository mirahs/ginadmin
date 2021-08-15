package page

import (
	"fmt"
	"ginadmin/model"
	"github.com/gin-gonic/gin"
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
	var wheres [][]interface{}
	return work(ctx, modelDatas, wheres, "")
}

func PageWhere(ctx *gin.Context, modelDatas interface{}, wheres [][]interface{}) *Info {
	return work(ctx, modelDatas, wheres, "")
}

func PageWhereOrder(ctx *gin.Context, modelDatas interface{}, wheres [][]interface{}, order string) *Info {
	return work(ctx, modelDatas, wheres, order)
}


func work(ctx *gin.Context, modelDatas interface{}, wheres [][]interface{}, order string) *Info {
	db := model.Db

	if len(wheres) > 0 {
		sqlFormat, args := whereBuild(wheres)
		db = db.Where(sqlFormat, args...)
	}

	if order != "" {
		db = db.Order(order)
	}

	page, limit := getPageLimit(ctx)

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

func whereBuild(wheres [][]interface{}) (sqlFormat string, args []interface{}) {
	for _, vals := range wheres {
		if len(vals) != 3 && len(vals) != 2 {
			panic(fmt.Errorf("error in query condition:%v", vals))
		}

		if sqlFormat != "" {
			sqlFormat += " AND "
		}

		switch len(vals) {
		case 2:
			sqlFormat += fmt.Sprintf("`%s`=?", vals[0])
			args = append(args, vals[1])
		case 3:
			switch vals[1] {
			case "in", "IN", "iN", "In":
				sqlFormat += fmt.Sprintf("`%s` IN (?)", vals[0])
			default:
				sqlFormat += fmt.Sprintf("`%s` %s ?", vals[0], vals[1])
			}
			args = append(args, vals[2])
		}
	}

	return sqlFormat, args
}

func getQuery(ctx *gin.Context) string {
	var datas []string
	query := ctx.Request.URL.Query()

	for key, val := range query {
		datas = append(datas, fmt.Sprintf("%s=%s", key, val[0]))
	}

	return "&" + strings.Join(datas, "&")
}
