package vo


type LogAdmUser struct {
	Account string `form:"account" validate:"min=1"`
}
