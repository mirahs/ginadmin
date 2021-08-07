package menu

import (
	"encoding/json"
)


// 菜单
type Menu struct {
	Code    string
	Name    string
	Data	[]interface{}
}

// 子菜单(二级菜单)
type MenuSub struct {
	Name 	string
	Data 	[]*MenuItem
}

// 菜单项
type MenuItem struct {
	Code	string
	Name    string
	Url     string
	Key     []int
}


var (
	menus []*Menu //菜单列表
	menuKeysStr string
)

// 获取菜单列表
func Menus() []*Menu {
	if menus != nil {
		return menus
	}

	menus = []*Menu{
		menuHome(),
		menuSys(),
	}

	return menus
}

func MenuKeysStr() string {
	if menuKeysStr != "" {
		return menuKeysStr
	}

	var keys []string
	menus := Menus()

	for _, menu := range menus {
		keys = append(keys, menu.Code)
	}

	keyBytes, _ := json.Marshal(keys)
	menuKeysStr = string(keyBytes)

	return menuKeysStr
}
