package main

import (
	"flag"
	"fmt"
	"ginadmin/conf"
	"ginadmin/model"
	"ginadmin/route"
	"ginadmin/thirdparty/pongo2gin"
	"ginadmin/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)


func main() {
	// 正式环境启动命令 ./ginadmin -mode prod
	mode := flag.String("mode", "dev", "运行环境 dev|test|prod")
	flag.Parse()

	conf.Load(*mode)
	model.Init()
	defer model.Close()
	util.IpInit(conf.App.Ip2RegionDbFile)
	defer util.IpClose()

	gin.SetMode(conf.App.GinMode)
	engine := gin.Default()

	// 模板引擎使用 pongo2
	engine.HTMLRender = pongo2gin.New(pongo2gin.RenderOptions{TemplateDir: conf.App.TemplateDir, ContentType: "text/html; charset=utf-8"})

	// 静态资源不是网络地址则 gin 负责提供静态资源访问
	if !strings.HasPrefix(conf.App.StaticUrl, "http") {
		engine.StaticFS(conf.App.StaticUrl, http.Dir(conf.App.StaticDir))
	}

	route.Init(engine)

	panic(engine.Run(fmt.Sprintf(":%d", conf.App.GinPort)))
}
