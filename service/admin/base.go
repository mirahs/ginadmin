package admin

import "ginadmin/repository"


// 所有 repository, service 直接访问, 控制器通过 service 间接访问
type base struct {
	RepoAdmUser         repository.AdmUser
	RepoLogAdmUserLogin repository.LogAdmUserLogin
}
