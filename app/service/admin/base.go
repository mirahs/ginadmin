package admin

import "ginadmin/app/repository"


// 所有 repository, service 直接访问(控制器间接访问)
type base struct {
	RepoAdmUser         repository.AdmUser
	RepoLogAdmUserLogin repository.LogAdmUserLogin
}
