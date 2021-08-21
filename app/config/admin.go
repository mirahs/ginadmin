package config


var (
	AdminUserTypeAdmin uint8 = 10
	AdminUserTypeGuest uint8 = 20

	AdminUserTypesDesc = map[uint8]string{
		AdminUserTypeAdmin: "管理员",
		AdminUserTypeGuest: "游客",
	}
)


// 获取用户类型名称
func GetTypeName(userType uint8) string {
	userName, _ := AdminUserTypesDesc[userType]
	return userName
}
