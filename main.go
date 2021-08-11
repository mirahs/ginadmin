package main

import (
	"ginadmin/model"
	"ginadmin/thirdparty/ip2region"
	"ginadmin/thirdparty/pongo2gin"
	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {
	model.DbInit()
	ip2region.RegionInit("./ip2region.db")
	defer ip2region.Ip.Close()

	engine := gin.Default()
	engine.HTMLRender = pongo2gin.Default()
	engine.StaticFS("/static", http.Dir("./static"))

	initRoutes(engine)

	panic(engine.Run())
}
