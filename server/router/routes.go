package router

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/keyslapperdev/task-manager-mono/server/models"
)

var validate = validator.New()

// add user routes
// create

func AddTask(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var task models.Task

	if err := c.BindJSON(&task); err != nil {
		c.JSON(returnBadRequest(err.Error()))
		fmt.Println(err)
		return
	}

	validationErr := validate.Struct(task)
	if validationErr != nil {
		c.JSON(returnBadRequest(validationErr.Error()))
		fmt.Println(validationErr)
		return
	}

	newTask := dataMgr.CreateTask(ctx, task)

	c.JSON(returnOK(newTask))
}

func GetTasks(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	tasks := dataMgr.GetTasks(ctx)

	c.JSON(returnOK(tasks))
}

func GetTaskByID(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	id, found := c.GetQuery("id")
	if !found {
		c.JSON(returnBadRequest("id not found in query"))
	}

	taskID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(returnBadRequest("bad id passed: " + id))
	}

	task := dataMgr.GetTaskByID(ctx, uint(taskID))

	c.JSON(returnOK(task))
}

func UpdateTask(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var task models.Task

	if err := c.BindJSON(&task); err != nil {
		c.JSON(returnBadRequest(err.Error()))
		fmt.Println(err)
		return
	}

	validationErr := validate.Struct(task)
	if validationErr != nil {
		c.JSON(returnBadRequest(validationErr.Error()))
		fmt.Println(validationErr)
		return
	}

	// the full task should be sent and returned, along with
	// the updated fields.
	updatedTask := dataMgr.UpdateTask(ctx, task)

	c.JSON(returnOK(updatedTask))
}

func DeleteTask(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var task models.Task

	if err := c.BindJSON(&task); err != nil {
		c.JSON(returnBadRequest(err.Error()))
		fmt.Println(err)
		return
	}

	validationErr := validate.Struct(task)
	if validationErr != nil {
		c.JSON(returnBadRequest(validationErr.Error()))
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

	c.JSON(returnOK(updatedTask))
}

func GetStatusMap(c *gin.Context) {
	statusMap := models.GetStatusMap()
	c.JSON(returnOK(statusMap))
}

type structuredReturn struct {
	Data   any      `json:"data,omitempty"`
	Errors []string `json:"errors,omitempty"`
}

func returnOK(data any) (int, structuredReturn) {
	return http.StatusOK, structuredReturn{Data: data}
}

func returnBadRequest(errs ...string) (int, structuredReturn) {
	return http.StatusBadRequest, structuredReturn{Errors: errs}
}
