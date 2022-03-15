package conf

import "gopkg.in/ini.v1"


type app struct {
	GinMode string //gin 模式(gin.DebugMode|gin.ReleaseMode|gin.TestMode)
	GinPort int

	MysqlHost     string
	MysqlPort     int
	MysqlDatabase string
	MysqlUser     string
	MysqlPassword string

	InitAccount     string //后台初始账号
	InitPassword    string //后台初始密码
	DefaultPassword string //后台新建账号默认密码

	StaticUrl string //静态资源地址 相对地址(默认 /static/) 由 gin 负责提供, 网络地址以 http 开头如 http://res.ginadmin.com/static/
	StaticDir string //静态资源目录 当 StaticUrl 为 相对地址 才需要配置并且生效

	TemplateDir string //模板目录

	SessionName   string //session 名称
	SessionSecret string //session 加密密钥

	Ip2RegionDbFile string //Ip2Region 数据文件

	UrlLogin  string //登录url
	UrlLogout string //登出url
	UrlIndex  string //主页url
	UrlDeny   string //访问拒绝url
}


var App *app


// 加载 app 配置文件
func loadApp(mode string)  {
	App = new(app)

	err := ini.MapTo(App, "./conf/app." + mode + ".ini")
	if err != nil {
		panic(err)
	}
}
