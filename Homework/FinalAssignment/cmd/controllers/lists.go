package controllers

import (
	"database/sql"
	"final/cmd/model"

	"github.com/gin-gonic/gin"
)

func GetLists() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*sql.DB)
		lists, err := model.NewRepository(db).GetLists()
		if err != nil {
			return
		}

		c.JSON(200, gin.H{
			"lists": lists,
		})
	}
}

func CreateList() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*sql.DB)
		var list model.List
		c.BindJSON(&list)
		list, err := model.NewRepository(db).CreateList(list.Name)
		if err != nil {
			return
		}

		c.JSON(200, gin.H{
			"list": list,
		})
	}
}

func DeleteList() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*sql.DB)
		var list model.List
		c.BindJSON(&list)
		err := model.NewRepository(db).DeleteList(list.ID)
		if err != nil {
			return
		}
		c.JSON(200, nil)
	}
}
