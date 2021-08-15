package menu

import (
	"ginadmin/config"
	"github.com/unknwon/com"
)

// 菜单
type Menu struct {
	Code string
	Name string
	Data []interface{}
}

// 子菜单(二级菜单)
type MenuSub struct {
	Name string
	Data []*MenuItem
}

// 菜单项
type MenuItem struct {
	Code string
	Name string
	Url  string
	Key  []int64
}

var (
	menus = []*Menu{
		menuHome(),
		menuSys(),
	}
	userTypeMenus = make(map[uint8][]*Menu) //用户类型对应菜单列表
)

// 初始化菜单数据
func init() {
	for userType, _ := range config.AdminUserTypesDesc {
		var userMenus []*Menu

		for _, menu := range menus {
			var datas []interface{}

			for _, data := range menu.Data {
				switch data.(type) {
				case *MenuItem:
					data2 := data.(*MenuItem)
					if com.IsSliceContainsInt64(data2.Key, int64(userType)) {
						menuItem := &MenuItem{
							Code: data2.Code,
							Name: data2.Name,
							Url: data2.Url,
						}
						datas = append(datas, menuItem)
					}
				case *MenuSub:
					var datasSub []*MenuItem

					data2 := data.(*MenuSub)
					for _, dataSub := range data2.Data {
						if com.IsSliceContainsInt64(dataSub.Key, int64(userType)) {
							datasSub = append(datasSub, dataSub)
						}
					}

					if len(datasSub) > 0 {
						menuSub := &MenuSub{
							Name: data2.Name,
							Data: datasSub,
						}

						datas = append(datas, menuSub)
					}
				}
			}

			if len(datas) > 0 {
				menu2 := &Menu{
					Code: menu.Code,
					Name: menu.Name,
					Data: datas,
				}
				userMenus = append(userMenus, menu2)
			}
		}

		userTypeMenus[userType] = userMenus
	}
}

func Get(userType uint8) []*Menu {
	return userTypeMenus[userType]
}
