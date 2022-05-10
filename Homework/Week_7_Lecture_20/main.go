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
	Data DrinkInstructions
	url  string
}

func NewCoctailBartender(url string) coctailBartender {
	//check if url is valid
	return coctailBartender{url: url}
}

func (c *coctailBartender) bartender(input string) (DrinksResponsePayload, error) {
	scrapeUrl := c.url + "?s=" + input
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

func (c *coctailBartender) Start() {
	//start the coctail bartender
	input := ""
	
	for input != "nothing" {
		fmt.Println("Enter the name of the drink you want to know the instructions for:")
		fmt.Scanln(&input)
		payload, err := c.bartender(input)
		if err != nil {
			fmt.Println("Drink not found")
		}
		fmt.Println(payload)
	}
}

func main() {
	//get user input and pass it to the coctail bartender
	c := NewCoctailBartender("https://www.thecocktaildb.com/api/json/v1/1/search.php")
	c.Start()

}
