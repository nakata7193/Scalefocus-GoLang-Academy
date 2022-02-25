package main

import (
	"fmt"
	"math/rand"
	"time"
)

func citiesAndPrices() ([]string, []int) {
	rand.Seed(time.Now().UnixMilli())
	cityChoices := []string{"Berlin", "Moscow", "Chicago", "Tokyo", "London"}
	dataPointCount := 100
	cities := make([]string, dataPointCount) //declaring slice cities
	for i := range cities {
		cities[i] = cityChoices[rand.Intn(len(cityChoices))]
	}
	prices := make([]int, dataPointCount) //declaring slice prices
	for i := range prices {
		prices[i] = rand.Intn(100)
	}
	return cities, prices
}

func groupSlices(cities []string, prices []int) map[string][]int {
	towns := make(map[string][]int)
	for i, t := range cities {
		towns[t] = append(towns[t], prices[i])
	}
	return towns
}

func main() {
    cities, prices := citiesAndPrices()
    towns := groupSlices(cities, prices)
    for cities, prices := range towns {
        fmt.Println(cities, prices)
    }
}
