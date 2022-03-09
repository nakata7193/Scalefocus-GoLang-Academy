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
			deck.cards = append(deck.cards, Card{cardVal, cardSuit})
		}
	}
	return deck
}

func (deck *Deck) Shuffle() *Deck {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.cards), func(i, j int) { deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i] })
	return deck
}

func (deck *Deck) Deal() (card *Card) {
	if len(deck.cards) > 0 {
		card = &deck.cards[0]
		deck.cards = deck.cards[1:]
		deck.size--
	}
	return card
}
