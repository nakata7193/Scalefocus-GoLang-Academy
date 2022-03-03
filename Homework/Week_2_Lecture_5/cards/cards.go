package cards

type cardVal = int

const (
	two cardVal = iota + 2
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	J
	D
	K
	A
)

type cardSuit = int

const (
	club cardSuit = iota + 1
	diamond
	heart
	spade
)


type Card struct{
	CardVal cardVal
	CardSuit cardSuit
}


