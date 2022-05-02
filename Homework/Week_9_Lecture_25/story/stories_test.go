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
	got := scraper.getTopStoriesIDs(10)
	want := ids[:10]

	//Assert
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}

func TestTopStories(t *testing.T) {

	ids := []int{1}

	router := http.NewServeMux()
	stories := []TopStory{
		{
			Title: "title1",
			Score: 10,
		},
	}
	router.Handle("/v0/topstories.json", handleTopStoriesIDs(ids))
	router.Handle("/v0/item/", handleTopStories(stories))
	mockServer := httptest.NewServer(router)

	//Act
	scraper := NewNewsScraper(mockServer.URL)
	got := scraper.GetTopStories(1)
	want := stories

	//Assert
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}
