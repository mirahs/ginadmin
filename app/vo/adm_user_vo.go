package vo


type AdmUserVo struct {
	Id uint32 `form:"id"` //Id
	Account string `form:"account"` //账号
	Password string `form:"password"` //密码
	Type     uint8  `form:"type"` //类型 10:管理员|20:游客

	IsLocked uint8 `form:"is_locked"` //是否被锁住 0:否|1:是
	Remark string `form:"remark"` //备注
}
