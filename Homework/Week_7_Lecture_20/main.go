package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Drink struct {
	Description string `json:"strInstructions"`
}

type Scraper struct {
	name string
	url  string
	Data Drink
}

func newScraper(url string, name string) Scraper {
	return Scraper{url: url, name: name}
}

func (sc *Scraper) Start() {
	req, err := http.NewRequest("GET", sc.url, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()

	for sc.name != "nothing" {
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			break
		}

		payload := sc.Data
		json.NewDecoder(resp.Body).Decode(&payload)
		time.Sleep(time.Second)
	}
}

func main() {
	var name string
	flag.StringVar(&name, "n", "", "the name of the coctail you want to drink")
	flag.Parse()
	sc := newScraper("https://www.thecocktaildb.com/api/json/v1/1/search.php", name)
	sc.Start()
	fmt.Println(sc.Data)
	if name == "" {
		flag.PrintDefaults()
	}
	fmt.Printf("Hello %s!\n", name)
}
