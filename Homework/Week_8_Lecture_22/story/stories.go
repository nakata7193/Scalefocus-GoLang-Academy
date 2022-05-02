package story

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type TopStory struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Score int    `json:"score"`
}

type TopStoriesPayload struct {
	TopStories []TopStory
}

type NewsScraper struct {
	url  string
	Data []TopStory
}

func NewNewsScraper(url string) *NewsScraper {
	return &NewsScraper{url: url}
}

func (n *NewsScraper) Top10Stories() []int {
	req, err := http.NewRequest("GET", n.url+"/v0/topstories.json", nil)
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
	return IDs
}

func (n *NewsScraper) GetTopStories(IDs []int) []TopStory {
	req, err := http.NewRequest("GET", n.url, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, id := range n.Top10Stories() {
		req.URL.Path = "/v0/item/" + fmt.Sprint(id) + ".json"
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		var topStory TopStory
		json.NewDecoder(resp.Body).Decode(&topStory)
		n.Data = append(n.Data, topStory)
	}
	return n.Data
}
