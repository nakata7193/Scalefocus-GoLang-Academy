package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//create a struct for coctail bartender

type DrinkInstructions struct {
	Recepy string `json:"strInstructions"`
}

type DrinksResponsePayload struct {
	Description DrinkInstructions
}

type coctailBartender struct {
	Data  DrinkInstructions
	url   string
	input string
}

func NewCoctailBartender(url string, input string) coctailBartender {
	//check if url is valid
	return coctailBartender{url: url, input: input}
}

func (c *coctailBartender) Start() (DrinksResponsePayload, error) {
	scrapeUrl := c.url + "?s=" + c.input
	req, err := http.NewRequest("GET", scrapeUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	//check the status code
	if res.StatusCode != 200 {
		return DrinksResponsePayload{}, err
	}
	defer res.Body.Close()

	payload := DrinksResponsePayload{}
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()
	return payload, nil
}

func main() {
	//get user input and pass it to the coctail bartender
	var input string
	fmt.Println("Enter the name of the drink you want to know the instructions for:")
	fmt.Scanln(&input)
	c := NewCoctailBartender("https://www.thecocktaildb.com/api/json/v1/1/search.php", input)
	for input != "nothing" {
		coctailBartender := NewCoctailBartender(c.url, c.input)
		coctailBartender.Start()
		fmt.Println(coctailBartender.Data)
		fmt.Scanln(&input)
	}
}

