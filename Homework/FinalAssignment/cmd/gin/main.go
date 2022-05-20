package main

import (
	"final/cmd"
	"final/cmd/controllers"
	"final/cmd/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

func main() {
	repository := utils.DbInit()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	//GET /api/lists/:id/tasks
	router.GET("/api/lists/:id/tasks", controllers.GetTasks(repository))

	//POST /api/lists/:id/tasks
	router.POST("/api/lists/:id/tasks", controllers.CreateTask(repository))

	//PATCH /tasks/:id
	router.PATCH("/api/tasks/:id", controllers.ToggleTask(repository))

	//DELETE /tasks/:id
	router.DELETE("/api/tasks/:id", controllers.DeleteTask(repository))

	// GET /api/lists/
	router.GET("/api/lists", controllers.GetLists(repository))

	//POST /api/lists/
	router.POST("/api/lists", controllers.CreateList(repository))

	//DELETE /api/lists/:id
	router.DELETE("/api/lists/:id", controllers.DeleteList(repository))

	//GET CSV file
	router.GET("/api/list/export", controllers.CSVExport(repository))

	//Weather API endpoint
	router.GET("/api/weather", controllers.GetWeather)

	// Do not touch this line!
	log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))
}
