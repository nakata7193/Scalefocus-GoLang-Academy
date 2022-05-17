package controllers

import (
	"final/cmd/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWeather(c *gin.Context) {
	var latitude, longitude string
	c.BindJSON(&longitude)
	c.BindJSON(&latitude)
	weather, err := utils.GetWeather(latitude, longitude)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	
	c.JSON(http.StatusOK, gin.H{
		"formatedTemp": weather["temp"],
		"description":  weather["description"],
		"city":         weather["city"],
	})
}
