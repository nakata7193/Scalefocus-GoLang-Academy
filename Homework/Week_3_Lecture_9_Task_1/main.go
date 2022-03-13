package main

import (
	"fmt"
	"log"

	carddraw "github.com/nakata7193/CardDraw"
	cardgame "github.com/nakata7193/CardGame"
)

func main() {
	deck := &cardgame.Deck{}
	deck.MakeDeck()
	deck.Shuffle()
	drawAllCards, err := carddraw.DrawAllCards(deck)
	if err != nil {
		log.Fatal(err)
		return
	}
	
	fmt.Print(drawAllCards)

}
