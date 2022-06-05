package vo


type AdmUserVo struct {
	Id uint32 `form:"id"` //Id
	Account string `form:"account"` //账号
	Password string `form:"password"` //密码
	Type     uint8  `form:"type"` //类型 10:管理员|20:游客

	IsLocked uint8 `form:"is_locked"` //是否被锁住 0:否|1:是
	Remark string `form:"remark"` //备注
}


type AdmUserDelOrReset struct {
	Id uint32 `form:"id"`
	IsLocked uint8 `form:"is_locked"`
}

type AdmUserAccount struct {
	Account string `form:"account"`
}

type AdmUserLogin struct {
	Account string `form:"account" validate:"min=1"`
	Password string `form:"password" validate:"min=1"`
}

type AdmUserPassword struct {
	Account string `form:"account" validate:"min=1"`
	Password string `form:"password" validate:"min=1"`
}

type AdmUserNew struct {
	Id uint32 `form:"id""`
	Account string `form:"account" validate:"min=1"`
	Type uint8 `form:"type" validate:"oneof=10 20"`
	Remark string `form:"remark" validate:"min=1"`
}
