package main

import (
	"fmt"
	"log"

	carddraw "github.com/nakata7193/CardDraw"
	cardgame "github.com/nakata7193/CardGame"
)

func main() {
	deck := &cardgame.Deck{}
	drawAllCards, err := carddraw.DrawAllCards(deck)
	deck.MakeDeck()
	deck.Shuffle()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(drawAllCards)

}
