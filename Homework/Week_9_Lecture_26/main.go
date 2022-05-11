package main

import (
	"database/sql"
	"log"
	"net/http"

	handler "github.com/nakata7193/handlers"
	repository "github.com/nakata7193/repository"
	_ "modernc.org/sqlite"
)

func main() {
	mux := http.NewServeMux()
	db, err := sql.Open("sqlite", "story.db")
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewRepository(db)
	mux.HandleFunc("/api/top", handler.TopStoriesHandler(repo))
	mux.HandleFunc("/top", handler.HTMLHandler())
	http.ListenAndServe(":8080", mux)
}
