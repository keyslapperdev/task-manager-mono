package routes

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/keyslapperdev/task-manager-mono/server/models"
	"github.com/keyslapperdev/task-manager-mono/server/storage"
)

var validate = validator.New()
var once = sync.Once{}
var dataMgr storage.DataMgr

func SetDataMgr(dm storage.DataMgr) {
	once.Do(func() {
		dataMgr = dm
	})
}

func AddTask(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var task models.Task

	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	validationErr := validate.Struct(task)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}

	taskID := dataMgr.CreateTask(ctx, task)

	c.JSON(http.StatusOK, gin.H{"id": taskID})
}

func GetTasks(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	tasks := dataMgr.GetTasks(ctx)

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func GetTaskByID(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	id, found := c.GetQuery("id")
	if !found {
		panic("id not found in query")
	}

	taskID, err := strconv.Atoi(id)
	if err != nil {
		panic("id conversion error: " + err.Error())
	}

	task := dataMgr.GetTaskByID(ctx, uint(taskID))

	c.JSON(http.StatusOK, gin.H{"task": task})
}

func UpdateTask(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var task models.Task

	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	validationErr := validate.Struct(task)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}

	// the full task should be sent and returned, along with
	// the updated fields.
	updatedTask := dataMgr.UpdateTask(ctx, task)

	c.JSON(http.StatusOK, gin.H{"task": updatedTask})
}

func DeleteTask(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var task models.Task

	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	validationErr := validate.Struct(task)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}

	var updatedTask models.Task
	shouldDelete := c.Query("delete")
	if shouldDelete == "delete" {
		dataMgr.DeleteTask(ctx, task)
	} else {
		updatedTask = dataMgr.CloseTask(ctx, task)
	}

	c.JSON(http.StatusOK, gin.H{"task": updatedTask})
}

func GetStatusMap(c *gin.Context) {
	statusMap := models.GetStatusMap()
	c.JSON(http.StatusOK, gin.H{"statusMap": statusMap})
}
