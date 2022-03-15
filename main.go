package main

import (
	"flag"
	"ginadmin/conf"
)


func main() {
	// 正式环境启动命令 ./ginadmin -mode prod
	mode := flag.String("mode", "dev", "运行环境 dev|test|prod")
	flag.Parse()

	conf.Load(*mode)

	Start()
}
