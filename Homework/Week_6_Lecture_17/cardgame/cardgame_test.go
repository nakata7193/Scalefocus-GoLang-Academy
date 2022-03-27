package cardgame

import (
	"testing"
)

func TestCompareCards(t *testing.T) {
	//Arrange
	Cards := []struct {
		cardOne Card
		cardTwo Card
		result  int
	}{
		{cardOne: Card{CardSuit: 2, CardVal: 7}, cardTwo: Card{CardSuit: 2, CardVal: 7}, result: 0},
		{cardOne: Card{CardSuit: 1, CardVal: 4}, cardTwo: Card{CardSuit: 2, CardVal: 7}, result: 1},
		{cardOne: Card{CardSuit: 2, CardVal: 7}, cardTwo: Card{CardSuit: 1, CardVal: 7}, result: -1},
	}

	for _, cardPair := range Cards {
		expectedResult := cardPair.result

		//Act
		funcResult := CompareCards(cardPair.cardOne, cardPair.cardTwo)

		//Assert
		if expectedResult != funcResult {
			t.Errorf("Expected %d, got %d", expectedResult, funcResult)
		}
	}
}

func TestMaxCard(t *testing.T) {
	//Arrange
	cardOne := Card{CardSuit: 2, CardVal: 7}
	cardTwo := Card{CardSuit: 4, CardVal: 7}
	cardThree := Card{CardSuit: 3, CardVal: 3}
	cardFour := Card{CardSuit: 4, CardVal: 2}
	Cards := []Card{cardOne, cardTwo, cardThree, cardFour}
	biggestCard := cardTwo

	//Act
	expectedMaxCard := MaxCard(Cards)

	//Assert
	if biggestCard != expectedMaxCard {
		t.Errorf("Expected %d, got %d", biggestCard, expectedMaxCard)
	}

}
