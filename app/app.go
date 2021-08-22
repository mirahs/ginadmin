package app

import (
	"fmt"
	"ginadmin/app/model"
	"ginadmin/app/thirdparty"
	"ginadmin/app/thirdparty/pongo2gin"
	"ginadmin/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)


// 启动 app
func Start() {
	model.DbInit()

	thirdparty.IpInit(config.App.Ip2RegionDbFile)
	defer thirdparty.IpClose()

	gin.SetMode(config.App.GinMode)
	engine := gin.Default()
	// 模板引擎使用 pongo2
	engine.HTMLRender = pongo2gin.New(pongo2gin.RenderOptions{TemplateDir: config.App.TemplateDir, ContentType: "text/html; charset=utf-8"})
	// 静态资源不是网络地址则 gin 负责提供静态资源访问
	if !strings.HasPrefix(config.App.StaticUrl, "http") {
		engine.StaticFS(config.App.StaticUrl, http.Dir(config.App.StaticDir))
	}

	initRoutes(engine)

	panic(engine.Run(fmt.Sprintf(":%d", config.App.GinPort)))
}
