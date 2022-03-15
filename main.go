package main

import (
	"ginadmin/app"
	"ginadmin/config"
)


func main() {
	config.Load("./app.ini")

	app.Start()
}
