package menu

import (
	"ginadmin/app/common"
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
	pathKeys = make(map[string][]int64)     //path对应的用户类型
)


// 初始化菜单数据
func init() {
	for userType, _ := range common.AdminUserTypesDesc {
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

					pathKeys[data2.Url] = data2.Key
				case *MenuSub:
					var datasSub []*MenuItem

					data2 := data.(*MenuSub)
					for _, dataSub := range data2.Data {
						if com.IsSliceContainsInt64(dataSub.Key, int64(userType)) {
							menuItem := &MenuItem{
								Code: dataSub.Code,
								Name: dataSub.Name,
								Url: dataSub.Url,
							}
							datasSub = append(datasSub, menuItem)
						}

						pathKeys[dataSub.Url] = dataSub.Key
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

// 获取用户类型对应的菜单数据
func Get(userType uint8) []*Menu {
	return userTypeMenus[userType]
}

// 检查路由能否访问
func Check(path string, userType uint8) bool {
	keys, ok := pathKeys[path]
	if !ok {
		return false
	}
	return com.IsSliceContainsInt64(keys, int64(userType))
}
