package main

import (
	"fmt"

	cardgame "github.com/nakata7193/cardgame"
)

func main() {
	cardOne := cardgame.Card{CardSuit: 2, CardVal: 7}
	cardTwo := cardgame.Card{CardSuit: 4, CardVal: 7}
	cardThree := cardgame.Card{CardSuit: 3, CardVal: 3}
	cardFour := cardgame.Card{CardSuit: 4, CardVal: 2}

	cards := []cardgame.Card{}
	cards = append(cards, cardOne)
	cards = append(cards, cardTwo)
	cards = append(cards, cardThree)
	cards = append(cards, cardFour)

	fmt.Print(cardgame.CompareCards(cardOne, cardTwo))
	fmt.Print(cardgame.MaxCard(cards))
}
