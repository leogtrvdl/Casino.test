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

type numbersRoue struct {
	value int
	color string
	tier  int
	ligne int
	demi  int
}
