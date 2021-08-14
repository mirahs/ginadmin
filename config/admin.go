package config


var (
	AdminUserTypeAdmin uint8 = 10
	AdminUserTypeGuest uint8 = 20

	AdminUserTypesDesc = map[uint8]string{
		AdminUserTypeAdmin: "管理员",
		AdminUserTypeGuest: "游客",
	}
)
