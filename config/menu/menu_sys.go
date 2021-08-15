package menu


// 菜单-管理
func menuSys() *Menu {
	menu := &Menu{
		Code: "sys",
		Name: "管理",
		Data: []interface{} {
			&MenuItem{
				Code: "password",
				Name: "密码更新",
				Url: "/admin/sys/password",
				Key: []int64{10, 20},
			},
			&MenuSub{
				Name: "管理员管理",
				Data: []*MenuItem{
					&MenuItem{
						Code: "master_new",
						Name: "添加管理员",
						Url: "/admin/sys/master_new",
						Key: []int64{10},
					},
					&MenuItem{
						Code: "master_list",
						Name: "管理员列表",
						Url: "/admin/sys/master_list",
						Key: []int64{10},
					},
				},
			},
			&MenuSub{
				Name: "日志",
				Data: []*MenuItem{
					&MenuItem{
						Code: "log_login",
						Name: "登录日志",
						Url: "/admin/sys/log_login",
						Key: []int64{10},
					},
				},
			},
		},
	}

	return menu
}
