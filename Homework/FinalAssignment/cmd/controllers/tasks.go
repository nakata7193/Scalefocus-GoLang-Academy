package controllers

import (
	"final/cmd/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

func GetTasks(data model.TaskOperations) gin.HandlerFunc {
	return func(c *gin.Context) {
		listID := c.Param("id")
		id, err := strconv.Atoi(listID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "func error"})
		}

		tasks, err := data.GetTasks(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid List"})
		}
		c.JSON(200, tasks)
	}
}

//Done: CreateTask
func CreateTask(data model.TaskOperations) gin.HandlerFunc {
	return func(c *gin.Context) {
		listID := c.Param("id")
		task := model.Task{}
		c.BindJSON(&task)
		id, err := strconv.Atoi(listID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "func error"})
		}
		err = data.CreateTask(id, task.Text)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Task"})
		}
		c.JSON(200, task)
	}
}

func ToggleTask(data model.TaskOperations) gin.HandlerFunc {
	return func(c *gin.Context) {
		taskID := c.Param("id")
		id, err := strconv.Atoi(taskID)

		data.ToggleTask(id)
		task := model.Task{}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "func error"})
		}

		c.BindJSON(&task)
		c.JSON(200, task)
	}
}

func DeleteTask(data model.TaskOperations) gin.HandlerFunc {
	return func(c *gin.Context) {
		taskID := c.Param("id")
		id, err := strconv.Atoi(taskID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "func error"})
		}

		err = data.DeleteTask(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Task"})
		}
	}
}