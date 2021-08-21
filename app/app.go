package app

import (
	"fmt"
	"ginadmin/app/config"
	"ginadmin/app/model"
	"ginadmin/app/thirdparty"
	"ginadmin/app/thirdparty/pongo2gin"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)


// 启动 app
func Start(app *config.App) {
	config.AppSet(app)

	model.DbInit()

	thirdparty.IpInit(config.AppInst.Ip2RegionDbFile)
	defer thirdparty.IpClose()

	gin.SetMode(app.GinMode)
	engine := gin.Default()
	// 模板引擎使用 pongo2
	engine.HTMLRender = pongo2gin.New(pongo2gin.RenderOptions{TemplateDir: config.AppInst.TemplateDir, ContentType: "text/html; charset=utf-8"})
	// 如果静态资源地址不以 http 开头则 gin 负责提供静态资源访问
	if !strings.HasPrefix(config.AppInst.StaticUrl, "http") {
		engine.StaticFS(config.AppInst.StaticUrl, http.Dir(config.AppInst.StaticDir))
	}

	initRoutes(engine)

	panic(engine.Run(fmt.Sprintf(":%d", config.AppInst.GinPort)))
}
