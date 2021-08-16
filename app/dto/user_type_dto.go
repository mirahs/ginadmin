package dto


type UserTypeDescDto struct {
	Type uint8
	Desc string
}


// 用户类型和名称是一个map, 转成切片的时候是无序的, 所以需要排序
type UTDSlice []*UserTypeDescDto

func (utd UTDSlice) Len() int {
	return len(utd)
}

func (utd UTDSlice) Less (i, j int) bool {
	return utd[i].Type < utd[j].Type
}

func (utd UTDSlice) Swap(i, j int) {
	utd[i], utd[j] = utd[j], utd[i]
}


func ToUserType(userType uint8, userDesc string) *UserTypeDescDto {
	return &UserTypeDescDto{
		Type: userType,
		Desc: userDesc,
	}
}
