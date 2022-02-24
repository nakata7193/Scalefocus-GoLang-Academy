package main

import (
	"fmt"
	"reflect"
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

cardOne := Card{4,3}

func compareCards(cardOne Card, cardTwo Card) int {
	//checks if suits are equal
	if reflect.DeepEqual(cardOne, cardTwo) {
		return 0
	} else if cardOne < cardTwo {
		return 1
	}
}

func main() {
	fmt.Print(compareCards())
}
