package controllers

import (
	"final/cmd/model"

	"github.com/gin-gonic/gin"
)

func GetLists(data model.ListOperations) gin.HandlerFunc {
	return func(c *gin.Context) {
		lists, err := data.GetLists()
		if err != nil {
			return
		}
		c.JSON(200, lists)
	}
}

func CreateList(data model.ListOperations) gin.HandlerFunc {
	return func(c *gin.Context) {
		var list model.List
		c.BindJSON(&list)
		list, err := data.CreateList(list)
		if err != nil {
			return
		}

		c.JSON(200, gin.H{
			"id":   list.ID,
			"name": list.Name,
		})
	}
}

func DeleteList(data model.ListOperations) gin.HandlerFunc {
	return func(c *gin.Context) {
		var list model.List
		c.BindJSON(&list)
		err := data.DeleteList(list)
		if err != nil {
			return
		}
	}
}
