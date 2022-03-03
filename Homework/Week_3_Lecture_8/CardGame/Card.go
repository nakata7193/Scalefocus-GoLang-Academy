package cardgame

type Card struct {
	CardVal  string
	CardSuit string
}

type Deck struct {
	Cards []Card
	size  int //just for Deal() test
}
