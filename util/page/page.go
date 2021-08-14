package page

import (
	"fmt"
	"ginadmin/model"
	"ginadmin/vm"
	"gorm.io/gorm"
)


type Info struct {
	Curr int
	Limit int
	Count int64
	Query string
	Datas interface{}
}


func Page(datasPtr interface{}, pageVm *vm.PageVm) *Info {
	var wheres map[string]interface{}
	return PageWhere(datasPtr, wheres, pageVm)
}

func PageWhere(datasPtr interface{}, wheres map[string]interface{}, pageVm *vm.PageVm) *Info {
	page, limit := fixPageLimit(pageVm.Page, pageVm.Limit)

	db := model.Db
	db = formatWhere(db, wheres)

	var count int64
	db.Model(datasPtr).Count(&count)

	db.Find(datasPtr)

	return &Info{
		Curr: page,
		Limit: limit,
		Count: count,
		Datas: datasPtr,
	}
}


func fixPageLimit(page, limit int) (int, int) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	return page, limit
}

func formatWhere(db *gorm.DB, wheres map[string]interface{}) *gorm.DB {
	for key, val := range wheres {
		db = db.Where(fmt.Sprintf("`%s`=?", key), val)
	}
	return db
}
