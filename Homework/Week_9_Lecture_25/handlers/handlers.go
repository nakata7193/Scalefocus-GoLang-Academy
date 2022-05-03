package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"time"

	stories "github.com/nakata7193/story"
)

type templateData struct {
	PageTitle string
	Links     []stories.TopStory
}

type Storage interface {
	GetLastTimeStamp() time.Time
	GetStory() []stories.TopStory
}

//create a handler for the top stories
func TopStoriesHandler(w http.ResponseWriter, r *http.Request, storage Storage) {
	scraper := stories.NewNewsScraper("https://hacker-news.firebaseio.com")
	threshold := time.Now().Add(-time.Hour)

	var storyList []stories.TopStory
	if storage.GetLastTimeStamp().Before(threshold) {
		storyList = scraper.GetTopStories(10)
	} else {
		storyList = storage.GetStory()
	} 
	
	json.NewEncoder(w).Encode(storyList)
}

//create html template handler for top stories
func HTMLHandler(w http.ResponseWriter, r *http.Request) {
	scraper := stories.NewNewsScraper("https://hacker-news.firebaseio.com")
	scraper.GetTopStories(10)
	tmpl := template.Must(template.ParseFiles("template.html"))
	data := templateData{
		PageTitle: "Top Stories",
		Links:     scraper.Data,
	}
	tmpl.Execute(w, data)
}
