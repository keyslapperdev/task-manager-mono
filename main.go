package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/keyslapperdev/task-manager-mono/server/routes"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.Default()
	router.Use(cors.Default())

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")
	// add user routes
	// create
	// delete
	api.POST("/task", routes.AddTask)
	api.GET("/tasks", routes.GetTasks)
	api.GET("/task", routes.GetTaskByID)
	api.PUT("/task", routes.UpdateTask)
	api.DELETE("/task", routes.DeleteTask)

	router.Run(":" + port)
}
