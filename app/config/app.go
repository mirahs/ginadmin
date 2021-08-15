package config


type App struct {
	MysqlHost     string //MySQL 主机
	MysqlPort     int    //MySQL 端口
	MysqlDatabase string //MySQL 数据库
	MysqlUser     string //MySQL 账号
	MysqlPassword string //MySQL 密码

	DefaultAccount     string //默认后台账号
	DefaultPassword    string //默认后台密码
	DefaultType        uint8  //默认后台账号类型(建议管理员, 登录后台创建新管理员账号后删除默认的后台账号)
	DefaultNewPassword string //新建后台账号默认密码

	TemplateDir string //模板目录
	StaticDir   string //静态资源目录

	SessionName   string //session 名称
	SessionSecret string //session 加密密钥

	Ip2RegionDbFile string //Ip2Region 数据文件

	UrlLogin	string //登录url
	UrlLogout	string //登录url
	UrlIndex	string //主页url
	UrlDeny		string //访问拒绝url
}


var AppInst *App


func AppSet(app *App) {
	AppInst = app
}
