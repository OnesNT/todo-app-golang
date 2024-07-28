package main

import (
	"log"
	"todo-app/config"
	"todo-app/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	r := gin.Default()
	config.ConnectDatabase()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
