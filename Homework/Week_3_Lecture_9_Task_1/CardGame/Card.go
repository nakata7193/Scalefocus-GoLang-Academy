package cardgame

type Card struct {
	cardVal  string
	cardSuit string
}

type Deck struct {
	cards []Card
	size  int  //just for Deal() test
}
