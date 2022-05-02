package story

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func handleTopStoriesIDs(ids []int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(ids)
	}
}

func handleTopStories(Stories []TopStory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id int
		var resultStory TopStory
		for _, story := range Stories {
			if story.ID == id {
				resultStory = story
			}
		}

		json.NewEncoder(w).Encode(resultStory)
	}
}

func TestTopStoriesIDs(t *testing.T) {
	router := http.NewServeMux()
	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	router.Handle("/v0/topstories.json", handleTopStoriesIDs(ids))
	mockServer := httptest.NewServer(router)

	//Act
	scraper := NewNewsScraper(mockServer.URL)
	got := scraper.Top10Stories()
	want := ids[:10]

	//Assert
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}

func TestTopStories(t *testing.T) {

	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	router := http.NewServeMux()
	stories := []TopStory{
		{ID: 1, Title: "title1", Score: 1},
		{ID: 2, Title: "title2", Score: 2},
		{ID: 3, Title: "title3", Score: 3},
		{ID: 4, Title: "title4", Score: 4},
		{ID: 5, Title: "title5", Score: 5},
		{ID: 6, Title: "title6", Score: 6},
		{ID: 7, Title: "title7", Score: 7},
		{ID: 8, Title: "title8", Score: 8},
		{ID: 9, Title: "title9", Score: 9},
		{ID: 10, Title: "title10", Score: 10},
		{ID: 11, Title: "title11", Score: 11},
	}

	router.Handle("/v0/item/", handleTopStories(stories))
	mockServer := httptest.NewServer(router)

	//Act
	scraper := NewNewsScraper(mockServer.URL)
	got := scraper.GetTopStories(ids)
	want := stories

	//Assert
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}
