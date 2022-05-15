package controllers

import (
	"final/cmd/model"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

func GetTasks(data model.TaskOperations) gin.HandlerFunc {
	return func(c *gin.Context) {
		var list model.List
		c.BindJSON(&list)
		tasks, err := data.GetTasks(list)
		if err != nil {
			return
		}
		c.JSON(200, tasks)
	}
}

func CreateTask(data model.TaskOperations) gin.HandlerFunc {
	return func(c *gin.Context) {
		var list model.List
		var task model.Task
		c.BindJSON(&list)
		c.BindJSON(&task)
		task, err := data.CreateTask(task, list)

		if err != nil {
			return
		}
		c.JSON(200, gin.H{
			"id":        task.ID,
			"text":      task.Text,
			"list":      list.ID,
			"completed": task.Completed,
		})
	}
}

func ToggleTask(data model.TaskOperations) gin.HandlerFunc {
	return func(c *gin.Context) {
		var task model.Task
		c.BindJSON(&task)
		task, err := data.ToggleTask(task)
		if err != nil {
			return
		}
		c.JSON(200, gin.H{
			"id":        task.ID,
			"text":      task.Text,
			"list":      task.ListID,
			"completed": task.Completed,
		})
	}
}

func DeleteTask(data model.TaskOperations) gin.HandlerFunc {
	return func(c *gin.Context) {
		var task model.Task
		c.BindJSON(&task)
		err := data.DeleteTask(task)
		if err != nil {
			return
		}
	}
}
