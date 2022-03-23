package main

import (
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type Metadata struct {
	URL  string
	Size int
}

func pingURL(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Printf("Got response for %s: %d\n", url, resp.StatusCode)
	return nil
}
func fetchURLs(urls []string, concurrency int) chan Metadata {
	processQueue := make(chan string, concurrency)
	outChan := make(chan Metadata)
	var wg sync.WaitGroup

	go func() {
		for _, urlToProcess := range urls {
			wg.Add(1)
			processQueue <- urlToProcess

			go func(url string) {
				defer wg.Done()
				time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
				log.Println("fetched: ", url)
				<-processQueue
				outChan <- Metadata{URL: url, Size: rand.Intn(2000)}
			}(urlToProcess)
		}
		wg.Wait()
		close(outChan)
	}()
	return outChan
}

func main() {
	urls := []string{
		"https://www.messenger.com/",
		"https://www.youtube.com/",
		"https://app.pluralsight.com/id",
		"https://hackforums.net/",
		"https://www.facebook.com/",
	}

	resultsChan := fetchURLs(urls, 2)
	for url := range resultsChan {
		log.Println("Done: ", url)
	}
}
