package app

import (
	"ginadmin/app/config"
	"ginadmin/app/model"
	"ginadmin/app/thirdparty"
	"ginadmin/app/thirdparty/pongo2gin"
	"github.com/gin-gonic/gin"
	"net/http"
)


// 启动app
func Start(app *config.App) {
	config.AppSet(app)

	model.DbInit()

	thirdparty.IpInit(config.AppInst.Ip2RegionDbFile)
	defer thirdparty.IpClose()

	gin.SetMode(app.GinMode)
	engine := gin.Default()
	// 模板引擎替换成 pongo2
	engine.HTMLRender = pongo2gin.New(pongo2gin.RenderOptions{TemplateDir: config.AppInst.TemplateDir, ContentType: "text/html; charset=utf-8"})
	engine.StaticFS("/static", http.Dir(config.AppInst.StaticDir))

	initRoutes(engine)

	panic(engine.Run())
}
