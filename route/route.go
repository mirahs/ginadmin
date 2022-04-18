package route

import (
	"ginadmin/api/admin"
	"ginadmin/conf"
	"ginadmin/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"net/http"
)


// 初始化路由
func Init(engine *gin.Engine) {
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
}


// 初始化 session
func initSession(group *gin.RouterGroup) {
	store := memstore.NewStore([]byte(conf.App.SessionSecret))
	store.Options(sessions.Options{Path: group.BasePath()})

	group.Use(sessions.Sessions(conf.App.SessionName, store))
}
