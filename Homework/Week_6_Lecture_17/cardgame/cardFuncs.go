package cardgame

func CompareCards(cardOne Card, cardTwo Card) int {
	//checks if suits are equal
	if cardOne.CardSuit == cardTwo.CardSuit {
		if cardOne.CardVal == cardTwo.CardVal {
			return 0
		} else if cardOne.CardVal < cardTwo.CardVal {
			return 1
		} else {
			return -1
		}
	} else if cardOne.CardSuit != cardTwo.CardSuit {
		if cardOne.CardSuit < cardTwo.CardSuit {
			return 1
		} else {
			return -1
		}
	} else {
		return -10
	}
}

func MaxCard(cards []Card) Card {

	var maxCard Card

	for i := range cards {
		if CompareCards(maxCard, cards[i]) == 1 {
			maxCard = cards[i]
		} else if CompareCards(maxCard, cards[i]) == -1 {
			continue
		} else {
			break
		}
	}
	return maxCard
}
