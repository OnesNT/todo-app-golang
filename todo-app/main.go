package main

import (
	"net/http"
	"todo-app/config"
	"todo-app/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	// Serve static files from the "public" directory
	r.Static("/static", "./public")

	// Serve index.html at the root
	r.GET("/", func(c *gin.Context) {
		c.File("./public/index.html")
	})

	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	config.ConnectDatabase()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
