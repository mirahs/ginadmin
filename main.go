package main

import (
	"ginadmin/pongo2gin"
	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {
	engine := gin.Default()

	engine.HTMLRender = pongo2gin.Default()

	engine.StaticFS("/static", http.Dir("./static"))

	initRoutes(engine)

	panic(engine.Run())
}
