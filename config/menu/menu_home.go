package menu


// 菜单-首页
func menuHome() *Menu {
	menu := &Menu{
		Code: "home",
		Name: "首页",
		Data: []interface{} {
			&MenuItem{
				Code: "welcome",
				Name: "欢迎",
				Url: "/admin/home/welcome",
				Key: []int64{10, 20},
			},
		},
	}

	return menu
}
