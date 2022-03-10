package carddraw

import (
	cardgame "github.com/nakata7193/CardGame"
)

type dealer interface {
	Deal() *cardgame.Card
}

func done(d dealer) (b bool) {
	if d.Deal() == nil {
		return false
	} else {
		return true
	}
}

func DrawAllCards(d dealer) (cards []cardgame.Card, err error) {
	for {
		//What err should I write??
		deckIsNotEmpty := done(d)
		if deckIsNotEmpty {
			cards = append(cards, *d.Deal())
		} else {
			return cards, err
		}
	}
}
