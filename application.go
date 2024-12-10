package main

import (
	"backend-todo/config"
	"backend-todo/routes"
)

func main() {
	config.Initialize()
	routes.Initialize()
}
