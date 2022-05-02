package main

import (
	"encoding/json"
	"html/template"
	"net/http"

	stories "github.com/nakata7193/story"
)

type templateData struct {
	PageTitle string
	Links     []stories.TopStory
}


//create a handler for the top stories
func TopStoriesHandler(w http.ResponseWriter, r *http.Request) {
	scraper := stories.NewNewsScraper("https://hacker-news.firebaseio.com")
	IDs := scraper.Top10Stories()
	scraper.GetTopStories(IDs)
	topStoriesPayload := stories.TopStoriesPayload{TopStories: scraper.Data}
	json.NewEncoder(w).Encode(topStoriesPayload)
}

//create html template handler for top stories
func HTMLHandler(w http.ResponseWriter, r *http.Request) {
	scraper := stories.NewNewsScraper("https://hacker-news.firebaseio.com")
	IDs := scraper.Top10Stories()
	scraper.GetTopStories(IDs)
	tmpl := template.Must(template.ParseFiles("template.html"))
	data := templateData{
		PageTitle: "Top Stories",
		Links:     scraper.Data,
	}
	tmpl.Execute(w, data)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/top", TopStoriesHandler)
	mux.HandleFunc("/top", HTMLHandler)
	http.ListenAndServe(":8080", mux)
}
