package router

import (
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/keyslapperdev/task-manager-mono/server/storage"
)

var dataMgr storage.DataMgr

func SetupRouter(mgr storage.DataMgr) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	dataMgr = mgr

	api := router.Group("/api")
	setTaskRoutes(api)
	setMiscRoutes(api)

	return router
}

func setTaskRoutes(rg *gin.RouterGroup) {
	rg.GET("/tasks", GetTasks)
	rg.GET("/task", GetTaskByID)
	rg.POST("/task", AddTask)
	rg.PATCH("/task", UpdateTask)
	rg.DELETE("/task", DeleteTask)
}

func setMiscRoutes(rg *gin.RouterGroup) {
	rg.GET("/statuses", GetStatusMap)
}
