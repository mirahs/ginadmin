package admin

import "ginadmin/app/service/admin"


// 所有 service, 控制器直接访问
var (
	serviceIndex admin.Index
	serviceSys   admin.Sys
)
