package controllers

import (
	"final/cmd/model"
	"net/http"
	"strconv"

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

//DO NOT TOUCH THIS
func CreateList(data model.ListOperations) gin.HandlerFunc {
	return func(c *gin.Context) {
		list := model.List{}
		err := data.CreateList(list.Name)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid List"})
		}
		c.JSON(200, list)
	}
}

func DeleteList(data model.ListOperations) gin.HandlerFunc {
	return func(c *gin.Context) {
		listID := c.Param("id")
		id, err := strconv.Atoi(listID)
		if err != nil {
			return
		}
		err = data.DeleteList(id)
		if err != nil {
			return
		}
	}
}
