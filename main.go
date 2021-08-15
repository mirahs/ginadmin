package main

import (
	"ginadmin/model"
	"ginadmin/thirdparty"
	"ginadmin/thirdparty/pongo2gin"
	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {
	model.DbInit()
	thirdparty.IpInit("./ip2region.db")
	defer thirdparty.Ip.Close()

	engine := gin.Default()
	engine.HTMLRender = pongo2gin.Default()
	engine.StaticFS("/static", http.Dir("./static"))

	initRoutes(engine)

	panic(engine.Run())
}
