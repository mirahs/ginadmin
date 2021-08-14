package vm


type PageVm struct {
	Page int `form:"page"` //当前页
	Limit int `form:"limit"` //显示数量
}
