package main

import (
	"database/sql"
	"final/cmd"
	"final/cmd/controllers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

type Repository struct {
	db *sql.DB
}

func main() {
	router := gin.Default()

	router.Use(func(ctx *gin.Context) {
		// This is a sample demonstration of how to attach middlewares in Gin
		log.Println("Gin middleware was called")
		ctx.Next()
	})

	//GET /api/lists/:id/tasks
	router.GET("/api/lists/:id/tasks", controllers.GetTasks())

	//POST /api/lists/:id/tasks
	router.POST("/api/lists/:id/tasks", controllers.CreateTask())

	//PATCH /tasks/:id
	router.PATCH("/api/tasks/:id", controllers.ToggleTask())

	//DELETE /tasks/:id
	router.DELETE("/api/tasks/:id", controllers.DeleteTask())

	// GET /api/lists/
	router.GET("/api/lists", controllers.GetLists())

	//POST /api/lists/
	router.POST("/api/lists", controllers.CreateList())

	//DELETE /api/lists/:id
	router.DELETE("/api/lists/:id", controllers.DeleteList())

	// Do not touch this line!
	log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))
}
