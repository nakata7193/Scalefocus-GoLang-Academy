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

func MaxCard(cards []Card) Card {

	var maxCard Card

	for i := range cards {
		if compareCards(maxCard, cards[i]) == 1 {
			maxCard = cards[i]
		} else if compareCards(maxCard, cards[i]) == -1 {
			continue
		} else {
			break
		}
	}
	return maxCard
}

func main() {
	cardOne := Card{CardSuit: 2, CardVal: 7}
	cardTwo := Card{CardSuit: 4, CardVal: 7}
	cardThree := Card{CardSuit: 3, CardVal: 3}
	cardFour := Card{CardSuit: 4, CardVal: 2}

	cards := []Card{}
	cards = append(cards, cardOne)
	cards = append(cards, cardTwo)
	cards = append(cards, cardThree)
	cards = append(cards, cardFour)

	fmt.Print(compareCards(cardOne, cardTwo))
	fmt.Print(MaxCard(cards))
}
