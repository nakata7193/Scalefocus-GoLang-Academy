package cardgame

import (
	"math/rand"
	"time"
)

func (deck *Deck) MakeDeck() *Deck {

	values := []string{"two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "J", "D", "K", "A"}
	suits := []string{"clubs", "diamonds", "hearts", "spades"}
	deck.size = 52
	for _, cardVal := range values {
		for _, cardSuit := range suits {
			deck.Cards = append(deck.Cards, Card{cardVal, cardSuit})
		}
	}
	return deck
}

func (deck *Deck) Shuffle() *Deck {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.Cards), func(i, j int) { deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i] })
	return deck
}

func (deck *Deck) Deal() (card *Card) {
	if deck.Cards == nil {
		return nil
	} else {
		card = &deck.Cards[0]
		deck.Cards = deck.Cards[1:]
		deck.size--
	}

	return card
}
