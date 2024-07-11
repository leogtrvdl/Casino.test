package games

type Player struct {
	Number int
	Name   string
	Hand   []Card
}

type Card struct {
	Value  string
	Symbol string
}

type NumbersRoue struct {
	Value int
	Color string
	Tier  int
	Ligne int
	Demi  int
}
