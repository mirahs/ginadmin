package main

import (
	"ginadmin/model"
	"ginadmin/thirdparty/pongo2gin"
	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {
	model.DbInit()

	engine := gin.Default()
	engine.HTMLRender = pongo2gin.Default()
	engine.StaticFS("/static", http.Dir("./static"))

	initRoutes(engine)

	panic(engine.Run())
}
