package main

import (
	"fmt"
	"ginadmin/conf"
	"ginadmin/model"
	"ginadmin/thirdparty"
	"ginadmin/thirdparty/pongo2gin"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)


// 启动 app
func Start() {
	model.DbInit()

	thirdparty.IpInit(conf.App.Ip2RegionDbFile)
	defer thirdparty.IpClose()

	gin.SetMode(conf.App.GinMode)
	engine := gin.Default()
	// 模板引擎使用 pongo2
	engine.HTMLRender = pongo2gin.New(pongo2gin.RenderOptions{TemplateDir: conf.App.TemplateDir, ContentType: "text/html; charset=utf-8"})
	// 静态资源不是网络地址则 gin 负责提供静态资源访问
	if !strings.HasPrefix(conf.App.StaticUrl, "http") {
		engine.StaticFS(conf.App.StaticUrl, http.Dir(conf.App.StaticDir))
	}

	initRoutes(engine)

	panic(engine.Run(fmt.Sprintf(":%d", conf.App.GinPort)))
}
