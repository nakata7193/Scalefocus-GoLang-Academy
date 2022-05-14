package controllers

import (
	"database/sql"
	"final/cmd/model"

	"github.com/gin-gonic/gin"
)

// tasks interface
type Task interface {
	GetTasks(listID int) ([]model.Task, error)
	CreateTask(listID int, taskName string) (model.Task, error)
	ToggleTask(taskID int) (model.Task, error)
	DeleteTask(taskID int) error
}

func GetTasks() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*sql.DB)
		var list model.List
		c.BindJSON(&list)
		taskList, err := model.NewRepository(db).GetTasks(list.ID)
		if err != nil {
			return
		}
		c.JSON(200, taskList)
	}
}

func CreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*sql.DB)
		var list model.List
		var task model.Task
		c.BindJSON(&list)
		c.BindJSON(&task)
		task, err := model.NewRepository(db).CreateTask(task.Text, list.ID)

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

//TODO: implement toggle task
func ToggleTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*sql.DB)
		var task model.Task
		c.BindJSON(&task)
		task, err := model.NewRepository(db).ToggleTask(task.ID)
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

func DeleteTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*sql.DB)
		var task model.Task
		c.BindJSON(&task)
		err := model.NewRepository(db).DeleteTask(task.ID)
		if err != nil {
			return
		}
		c.JSON(200, nil)
	}
}
