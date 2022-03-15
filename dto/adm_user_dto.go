package dto

import (
	"ginadmin/common"
	"ginadmin/model"
	"ginadmin/util"
)


type AdmUserDto struct {
	Id       uint32
	Account  string
	Type     uint8
	TypeDesc string

	IsLocked uint8
	IsLockedDesc string
	Remark string

	LoginTimes uint32
	LoginTime uint32
	LoginTimeDesc string
	LoginIp string
}


func ToAdmUser(user *model.AdmUser) *AdmUserDto {
	return &AdmUserDto{
		Id:       user.Id,
		Account:  user.Account,
		Type:     user.Type,
		TypeDesc: common.AdminUserTypesDesc[user.Type],

		IsLocked:     user.IsLocked,
		IsLockedDesc: util.If(user.IsLocked == 0, "锁住", "解锁").(string),
		Remark:       user.Remark,

		LoginTimes:    user.LoginTimes,
		LoginTime:     user.LoginTime,
		LoginTimeDesc: util.Time2Datetime(int64(user.LoginTime)),
		LoginIp:       user.LoginIp,
	}
}
