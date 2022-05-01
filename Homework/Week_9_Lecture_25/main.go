package main

import (
	"net/http"
	"github.com/nakata7193/stories"
	//IMPORT DRIVER
	_ "github.com/nakata7193/stories/driver/sqlite3"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/top", stories.TopStoriesHandler)
	mux.HandleFunc("/top", 	stories.HTMLHandler)
	http.ListenAndServe(":8080", mux)
}
