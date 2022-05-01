package stories

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type TopStories struct {
	Title string `json:"title"`
	Score int    `json:"score"`
}

type TopStoriesPayload struct {
	TopStories []TopStories
}

type NewsScraper struct {
	url  string
	Data []TopStories
}

type templateData struct {
	PageTitle string
	Links     []TopStories
}

func NewNewsScraper(url string) *NewsScraper {
	return &NewsScraper{url: url}
}

func Top10Stories() []string {
	req, err := http.NewRequest("GET", "https://hacker-news.firebaseio.com/v0/topstories.json", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	var IDs []int
	json.NewDecoder(resp.Body).Decode(&IDs)
	IDs = IDs[:10]
	var IDsString []string
	for _, id := range IDs {
		IDsString = append(IDsString, strconv.Itoa(id))
	}
	return IDsString
}

func (n *NewsScraper) GetTopStories() {
	req, err := http.NewRequest("GET", n.url, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, id := range Top10Stories() {
		req.URL.Path = "/v0/item/" + id + ".json"
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		var topStory TopStories
		json.NewDecoder(resp.Body).Decode(&topStory)
		n.Data = append(n.Data, topStory)
	}
}

//create a handler for the top stories
func TopStoriesHandler(w http.ResponseWriter, r *http.Request) {
	scraper := NewNewsScraper("https://hacker-news.firebaseio.com")
	scraper.GetTopStories()
	topStoriesPayload := TopStoriesPayload{TopStories: scraper.Data}
	json.NewEncoder(w).Encode(topStoriesPayload)
}

//create html template handler for top stories
func HTMLHandler(w http.ResponseWriter, r *http.Request) {
	scraper := NewNewsScraper("https://hacker-news.firebaseio.com")
	scraper.GetTopStories()
	tmpl := template.Must(template.ParseFiles("template.html"))
	data := templateData{
		PageTitle: "Top Stories",
		Links:     scraper.Data,
	}
	tmpl.Execute(w, data)
}
