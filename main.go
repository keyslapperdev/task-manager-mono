package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/keyslapperdev/task-manager-mono/server/routes"
	"github.com/keyslapperdev/task-manager-mono/server/storage"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.Default()
	router.Use(cors.Default())

	routes.SetDataMgr(storage.NewDBStorer(true))

	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")
	// add user routes
	// create
	api.POST("/task", routes.AddTask)
	api.GET("/tasks", routes.GetTasks)
	api.GET("/task", routes.GetTaskByID)
	api.GET("/statuses", routes.GetStatusMap)
	api.PATCH("/task", routes.UpdateTask)
	api.DELETE("/task", routes.DeleteTask)

	router.Run(":" + port)
}
