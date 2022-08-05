package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/keyslapperdev/task-manager-mono/server/models"
)

var validate = validator.New()

func AddTask(c *gin.Context) {
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
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

	// save task
	// return task id

	c.JSON(http.StatusOK, nil)
}

func GetTasks(c *gin.Context) {
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	// get tasks
	// return tasks

	c.JSON(http.StatusOK, nil)
}

func GetTaskByID(c *gin.Context) {
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	// get task by ID
	// return tasks

	c.JSON(http.StatusOK, nil)
}

func UpdateTask(c *gin.Context) {
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
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

	// update task
	// return task id

	c.JSON(http.StatusOK, nil)
}

func DeleteTask(c *gin.Context) {
	var _, cancel = context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Delete task
	// return task id

	c.JSON(http.StatusOK, nil)
}
