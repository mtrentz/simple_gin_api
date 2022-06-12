package main

import (
	"github.com/mtrentz/simple_gin_api/database"
	"github.com/mtrentz/simple_gin_api/routes"
)

func main() {
	database.Connect()
	routes.HandleRequest()
}