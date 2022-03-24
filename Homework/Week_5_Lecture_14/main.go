package main

import (
	"log"
	"net/http"
	"sync"
)

type Metadata struct {
	URL     string
	Message string
	Error   error
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
				err := pingURL(url)
				if err != nil {
					outChan <- Metadata{Error: err}
				} else {
					<-processQueue
					outChan <- Metadata{URL: url, Message: "ok", Error: nil}
				}
			}(urlToProcess)
		}
		wg.Wait()
		close(outChan)
	}()
	return outChan
}

func main() {

	/*Just import the CLI*/

	// urls := []string{
	// 	"https://www.messenger.com/",
	// 	"https://www.youtube.com/",
	// 	"https://app.pluralsight.com/id",
	// 	"https://hackforums.net/",
	// 	"https://www.facebook.com/",
	// }

	// resultsChan := fetchURLs(urls, 2)
	// for url := range resultsChan {
	// 	log.Println("Done: ", url)
	// }
}
