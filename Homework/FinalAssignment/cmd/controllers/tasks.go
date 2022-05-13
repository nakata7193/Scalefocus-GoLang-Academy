package controllers

import (
	"final/cmd/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController interface {
	GetTasks(c *gin.Context) error
	CreateTask(c *gin.Context) error
	ToggleTask(c *gin.Context) error
	DeleteTask(c *gin.Context) error
}

func CreateTask(c *gin.Context) error {
	var task model.Task
	c.BindJSON(&task)
	task, err := model.CreateTask(task.Text, task.ListID)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        task.ID,
		"text":      task.Text,
		"listId":    task.ListID,
		"completed": task.Completed,
	})
}
