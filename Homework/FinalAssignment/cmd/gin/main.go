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

	router := gin.Default()

	router.Use(func(ctx *gin.Context) {
		// This is a sample demonstration of how to attach middlewares in Gin
		log.Println("Gin middleware was called")
		ctx.Next()
	})

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

	// Do not touch this line!
	log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))
}
