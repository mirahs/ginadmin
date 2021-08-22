package main

import (
	"ginadmin/app"
	"ginadmin/config"
)


func main() {
	config.Load("./public/app.ini")

	app.Start()
}
