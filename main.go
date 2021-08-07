package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {
	engine := gin.Default()
	engine.StaticFS("/static", http.Dir("./static"))
	engine.LoadHTMLGlob("template/**/**/*")
	registerRoutes(engine)
	panic(engine.Run())
}
