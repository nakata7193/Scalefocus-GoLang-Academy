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
	if cardOneSuit != cardTwoSuit {
		if cardOneVal == cardTwoVal {
			return 0
		} else if cardOneVal > cardTwoVal{
			return -1
		} else{
			return 1
		}
	} 
	}
}//some kind of error here???

func main(){
	fmt.printf(compareCards(2,3,4,5))
}

