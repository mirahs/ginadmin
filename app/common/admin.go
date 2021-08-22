package common


const (
	AdminUserTypeAdmin uint8 = 10
	AdminUserTypeGuest uint8 = 20
)

var (
	AdminUserTypesDesc = map[uint8]string{
		AdminUserTypeAdmin: "管理员",
		AdminUserTypeGuest: "游客",
	}
)
