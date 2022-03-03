package carddraw

import cardgame "github.com/nakata7193/CardGame"

type dealer interface {
	Deal() *cardgame.Card
}

func DrawAllCards(dealer dealer) []cardgame.Card {

}
