package carddraw

import (
	cardgame "github.com/nakata7193/CardGame"
)

type dealer interface {
	Deal() *cardgame.Card
	Done() bool
}

func DrawAllCards(d dealer) (cards []cardgame.Card, err error) {
	for {
		deckIsEmpty := dealer.Done(d)
		if !deckIsEmpty {
			cards = append(cards, *d.Deal())
		} else {
			return cards, err
		}
	}
}
