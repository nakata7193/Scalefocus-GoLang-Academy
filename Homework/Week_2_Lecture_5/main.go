package main

import (
	"fmt"
)

type cardVal = int

const (
	two cardVal = iota + 2
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	J
	D
	K
	A
)

type cardSuit = int

const (
	club cardSuit = iota + 1
	diamond
	heart
	spade
)

type Card struct {
	CardVal  cardVal
	CardSuit cardSuit
}

func compareCards(cardOne Card, cardTwo Card) int {
	//checks if suits are equal
	if cardOne.CardSuit == cardTwo.CardSuit {
		if cardOne.CardVal == cardTwo.CardVal {
			return 0
		} else if cardOne.CardVal < cardTwo.CardVal {
			return 1
		} else {
			return -1
		}
	}
	//checks if suits are not equal
	if cardOne.CardSuit != cardTwo.CardSuit {
		if cardOne.CardSuit < cardTwo.CardSuit {
			return 1
		} else {
			return -1
		}
	}
	return -10
}

func maxCard(cards []Card) Card {

	var MaxCard Card

	for i := 0; i < len(cards); i++ {
		if compareCards(cards[i], cards[i+1]) == -1 {
			MaxCard = cards[i+1]
		} else if compareCards(cards[i], cards[i+1]) == 1 {
			MaxCard = cards[i]
		} else {
			MaxCard = cards[i]
		}
	}
	return MaxCard
}

func main() {
	cardOne := Card{CardSuit: 7, CardVal: 1}
	cardTwo := Card{CardSuit: 7, CardVal: 2}
	cardThree := Card{CardSuit: 9, CardVal: 3}
	cardFour := Card{CardSuit: 10, CardVal: 4}

	cards := []Card{}
	cards = append(cards, cardOne)
	cards = append(cards, cardTwo)
	cards = append(cards, cardThree)
	cards = append(cards, cardFour)

	fmt.Print(compareCards(cardOne, cardTwo))
	fmt.Print(MaxCard(cards))
}
