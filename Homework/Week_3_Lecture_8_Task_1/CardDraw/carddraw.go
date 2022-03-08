package carddraw

import (
	cardgame "github.com/nakata7193/CardGame"
)

type dealer interface {
	Deal() *cardgame.Card
}

func DrawAllCards(d dealer) (cards []cardgame.Card) {
	for d.Deal() != nil {
		cards = append(cards, *d.Deal())
	}
	return cards
}
