package routes

import (
	"todo-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/tasks", controllers.CreateTask)
	router.GET("/tasks", controllers.GetTasks)
	router.GET("/tasks/:id", controllers.GetTask)
	router.PUT("/tasks/:id", controllers.UpdateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)
}
