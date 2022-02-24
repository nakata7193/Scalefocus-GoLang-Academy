package main

import "github.com/nakata7193/Week_2_Lecture_5/cards"
import "fmt"

func compareCards(cardOne Card, cardTwo Card) int {
	//checks if suits are equal
	if cardOne == cardTwo {
		if cardOne == cardTwo {
			return 0
		} else if cardOne < cardTwo {
			return 1
		} else {
			return -1
		}
	}
	//checks if suits are not equal
	if cardOne != cardTwo {
		if cardOne < cardTwo {
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

