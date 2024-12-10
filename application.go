package main

import (
	"backend-todo/config"
	"backend-todo/routes"
	"backend-todo/services"
)

func main() {
	config.Initialize()
	services.Initialize()
	routes.Initialize()
}
