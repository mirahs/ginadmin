package main

import (
	"ginadmin/config"
)


func main() {
	config.Load("./app.ini")

	Start()
}
