package main

import (
	"flag"
	"fmt"
	"ginadmin/conf"
	"ginadmin/model"
	"ginadmin/route"
	"ginadmin/util"
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

	engine := route.Init()

	panic(engine.Run(fmt.Sprintf(":%d", conf.App.GinPort)))
}
