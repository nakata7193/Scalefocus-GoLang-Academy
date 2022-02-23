package main

import "fmt"

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

func compareCards(cardOneVal cardVal, cardOneSuit cardSuit, cardTwoVal cardVal, cardTwoSuit cardSuit) int {
	//checks if suits are equal
	if cardOneSuit == cardTwoSuit {
		if cardOneVal == cardTwoVal {
			return 0
		} else if cardOneVal < cardTwoVal {
			return 1
		} else {
			return -1 
		}
	}
	//checks if suits are not equal
	if cardOneSuit != cardTwoSuit{
		if cardOneSuit < cardTwoSuit{
			return 1
		} else {
			return -1
		}
	}
	return -10
}

func main() {
	fmt.Print(compareCards(7, 1, 7, 2))
}
