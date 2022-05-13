package main

import (
	"final/cmd"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

func main() {

	router := gin.Default()

	router.Use(func(ctx *gin.Context) {
		// This is a sample demonstration of how to attach middlewares in Gin
		log.Println("Gin middleware was called")
		ctx.Next()
	})

	//Tasks
	//GET /api/lists/:id/tasks
	router.GET("/api/lists/:id/tasks", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Successful",
		})
	})

	//POST /api/lists/:id/tasks
	router.POST("/api/lists/:id/tasks", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"text": "string",
		})
	})

	//PATCH /tasks/:id
	router.PATCH("/api/tasks/:id", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"completed": true,
		})
	})

	//DELETE /tasks/:id

	router.DELETE("/api/tasks/:id", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"id": 0,
		})
	})

	// Lists
	// GET /api/lists/
	router.GET("/api/lists", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Successful",
		})
	})

	//POST /api/lists/
	router.POST("/api/lists", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"id":   0,
			"name": "string",
		})
	})

	//DELETE /api/lists/:id
	router.DELETE("/api/lists/:id", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"id": 0,
		})
	})

	// Do not touch this line!
	log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))
}
