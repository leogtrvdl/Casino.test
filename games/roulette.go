package games

import (
	"fmt"
	"math/rand"
	"time"
)

func PlayRoulette() {
	numRoulette()
}

func roue() []numbersRoue {
	values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36}

	colors := []string{"red", "black", "green"}

	tiers := []int{1, 2, 3}

	var numRoulette []numbersRoue

	for _, value := range values {

		//couleur
		color := colors[0]
		switch {
		case value == 0:
			color = colors[2]
		case value%2 == 0:
			color = colors[1]
		}

		//tier
		tier := 0
		for i, t := range tiers {
			if value >= (t-1)*12+1 && value <= t*12 {
				tier = i + 1
				break
			}
		}

		//ligne
		ligne := 0
		switch {
		case value%3 == 1:
			ligne = 1
		case value%3 == 2:
			ligne = 2
		case value%3 == 0 && value != 0:
			ligne = 3
		}

		//demi
		demi := 0
		if value >= 1 && value <= 18 {
			demi = 1
		} else if value >= 19 && value <= 36 {
			demi = 2
		}

		roue := numbersRoue{value, color, tier, ligne, demi}
		numRoulette = append(numRoulette, roue)
	}
	return numRoulette
}

func numRoulette() {
	rand.Seed(time.Now().UnixNano())
	numeros := roue()
	randomIndex := rand.Intn(len(numeros))
	numero := numeros[randomIndex]
	fmt.Printf("Numéro: %d, Couleur: %s, Tier: %d, Ligne: %d, Demi: %d\n", numero.value, numero.color, numero.tier, numero.ligne, numero.demi)
}
