package config


type App struct {
	MysqlHost     string //MySQL 主机
	MysqlPort     int    //MySQL 端口
	MysqlDatabase string //MySQL 数据库
	MysqlUser     string //MySQL 账号
	MysqlPassword string //MySQL 密码

	InitAccount     string //初始后台管理员账号
	InitPassword    string //初始后台管理员密码
	DefaultPassword string //新建后台账号默认密码

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


var AppInst = &App{
	MysqlHost:     "127.0.0.1",
	MysqlPort:     3306,
	MysqlDatabase: "ginadmin",
	MysqlUser:     "root",
	MysqlPassword: "root",

	InitAccount:     "admin",
	InitPassword:    "admin",
	DefaultPassword: "123456",

	TemplateDir: "./app/template/",
	StaticDir:   "./app/static/",

	SessionName: "ginadmin_session",
	SessionSecret: "ginadmin_secret",

	Ip2RegionDbFile: "./app/ip2region.db",

	UrlLogin: "/admin/index/login",
	UrlLogout: "/admin/index/logout",
	UrlIndex: "/admin/index/index",
	UrlDeny: "/admin/index/deny",
}


// app参数设置(非零值替换)
func AppSet(app *App) {
	if app == nil {
		return
	}

	if app.MysqlHost != "" {
		AppInst.MysqlHost = app.MysqlHost
	}
	if app.MysqlPort != 0 {
		AppInst.MysqlPort = app.MysqlPort
	}
	if app.MysqlDatabase != "" {
		AppInst.MysqlDatabase = app.MysqlDatabase
	}
	if app.MysqlUser != "" {
		AppInst.MysqlUser = app.MysqlUser
	}
	if app.MysqlPassword != "" {
		AppInst.MysqlPassword = app.MysqlPassword
	}

	if app.InitAccount != "" {
		AppInst.InitAccount = app.InitAccount
	}
	if app.InitPassword != "" {
		AppInst.InitPassword = app.InitPassword
	}
	if app.DefaultPassword != "" {
		AppInst.DefaultPassword = app.DefaultPassword
	}

	if app.TemplateDir != "" {
		AppInst.TemplateDir = app.TemplateDir
	}
	if app.StaticDir != "" {
		AppInst.StaticDir = app.StaticDir
	}

	if app.SessionName != "" {
		AppInst.SessionName = app.SessionName
	}
	if app.SessionSecret != "" {
		AppInst.SessionSecret = app.SessionSecret
	}

	if app.Ip2RegionDbFile != "" {
		AppInst.Ip2RegionDbFile = app.Ip2RegionDbFile
	}

	if app.UrlLogin != "" {
		AppInst.UrlLogin = app.UrlLogin
	}
	if app.UrlLogout != "" {
		AppInst.UrlLogout = app.UrlLogout
	}
	if app.UrlIndex != "" {
		AppInst.UrlIndex = app.UrlIndex
	}
	if app.UrlDeny != "" {
		AppInst.UrlDeny = app.UrlDeny
	}
}
