package carddraw

import (
	cardgame "github.com/nakata7193/CardGame"
)

type dealer interface {
	Deal() *cardgame.Card
}

func DrawAllCards(d dealer) (cards []cardgame.Card) {
	for {
		i := d.Deal()
		if i != nil {
			cards = append(cards, *i)
		} else {
			return cards
		}
	}
}
