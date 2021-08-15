package dto


type UserTypeDescDto struct {
	Type uint8
	Desc string
}


func ToUserType(userType uint8, userDesc string) *UserTypeDescDto {
	return &UserTypeDescDto{
		Type: userType,
		Desc: userDesc,
	}
}
