package main

import (
	"net/http"

	_ "modernc.org/sqlite"
	handler "github.com/nakata7193/handlers"
	repository "github.com/nakata7193/repository"
)

func main() {
	mux := http.NewServeMux()
	repo := repository.NewRepository(nil)
	mux.HandleFunc("/api/top", handler.TopStoriesHandler(repo))
	mux.HandleFunc("/top", handler.HTMLHandler)
	http.ListenAndServe(":8080", mux)
}
