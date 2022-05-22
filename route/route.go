package route

import (
	"ginadmin/api/admin"
	"ginadmin/conf"
	"ginadmin/middleware"
	"ginadmin/thirdparty/pongo2gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)


// Init 初始化路由
func Init() *gin.Engine {
	engine := initEngine()

	engine.GET("/", func(context *gin.Context) {
		context.Redirect(http.StatusPermanentRedirect, conf.App.UrlIndex)
	})

	adminG := engine.Group("/admin")
	{
		initSession(adminG)

		adminG.Use(middleware.AdminValidate())

		adminG.GET("/", admin.Index)
		adminG.GET("/login", admin.Login)
		adminG.POST("/login", admin.Login)
		adminG.GET("/logout", admin.Logout)
		adminG.GET("/deny", admin.Deny)

		adminG.GET("/home/welcome", admin.HomeWelcome)

		adminG.GET("/sys/password", admin.SysPassword)
		adminG.POST("/sys/password", admin.SysPassword)
		adminG.GET("/sys/master_new", admin.SysMasterNew)
		adminG.POST("/sys/master_new", admin.SysMasterNew)
		adminG.GET("/sys/master_list", admin.SysMasterList)
		adminG.GET("/sys/log_login", admin.SysLogLogin)
	}

	return engine
}


func initEngine() *gin.Engine {
	gin.SetMode(conf.App.GinMode)

	var engine *gin.Engine
	if conf.App.GinMode == gin.DebugMode {
		engine = gin.Default()
	} else {
		engine = gin.New()
		engine.Use(gin.Recovery())
	}

	// 模板引擎使用 pongo2
	engine.HTMLRender = pongo2gin.New(pongo2gin.RenderOptions{TemplateDir: conf.App.TemplateDir, ContentType: "text/html; charset=utf-8"})

	// 静态资源不是网络地址则 gin 负责提供静态资源访问
	if !strings.HasPrefix(conf.App.StaticUrl, "http") {
		engine.StaticFS(conf.App.StaticUrl, http.Dir(conf.App.StaticDir))
	}

	return engine
}

// 初始化 session
func initSession(group *gin.RouterGroup) {
	store := memstore.NewStore([]byte(conf.App.SessionSecret))
	store.Options(sessions.Options{Path: group.BasePath()})

	group.Use(sessions.Sessions(conf.App.SessionName, store))
}
