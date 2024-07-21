package main

import (
	"todo-app/config"
	"todo-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()
	routes.SetupRoutes(r)
	r.Run()
}
