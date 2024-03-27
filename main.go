package main

import (
	"go_restaurant/config"
	"go_restaurant/routes"
)

func main() {
	config.InitDataBase()
	config.Migration()

	e := routes.Init()

	e.Start(":8080")
}