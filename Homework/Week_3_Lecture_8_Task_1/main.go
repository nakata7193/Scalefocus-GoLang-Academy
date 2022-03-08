package main

import (
	"fmt"
	carddraw "github.com/nakata7193/CardDraw"
	cardgame "github.com/nakata7193/CardGame"
)

func main() {
	deck := &cardgame.Deck{}
	deck.MakeDeck()
	deck.Shuffle()
	fmt.Print(carddraw.DrawAllCards(deck))

}
