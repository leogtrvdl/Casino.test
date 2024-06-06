package games

type Player struct {
	number int
	name   string
	hand   []Card
}

type Card struct {
	value  string
	symbol string
}

type NumbersRoue struct {
	Value int
	Color string
	Tier  int
	Ligne int
	Demi  int
}
