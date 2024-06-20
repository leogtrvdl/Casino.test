package games

import (
	"fmt"
	"strconv"
)

func PlayBlackjack() {

	//On commence le jeu

	//Affichage de la main du croupier
	fmt.Println("Main du croupier:")
	for _, card := range croupier.hand {
		fmt.Printf(" - %s de %s\n", card.value, card.symbol)
	}

	//Tour des joueurs
	for i := range joueurs {
		for {
			joueurs[i].PrintDetails()
			total := CalculateHand(joueurs[i].hand)
			fmt.Printf("Total des points : %d\n", total)

			if total >= 21 {
				break
			}

			var action string
			fmt.Println("Voulez vous tirer une autre carte? (oui/non)")
			fmt.Scanln(&action)

			if action == "oui" {
				GiveRandomCard(&joueurs[i], &deck)
			} else {
				break
			}
		}
	}

	//Tour du croupier
	for CalculateHand(croupier.hand) < 17 {
		GiveRandomCard(&croupier, &deck)
	}

	//Afficher a nouveau la main du croupier
	fmt.Println("Main du croupier:")
	for _, card := range croupier.hand {
		fmt.Printf(" - %s de %s\n", card.value, card.symbol)
	}

	//Comparaison des mains pour déclarer les gagnants
	croupierTotal := CalculateHand(croupier.hand)
	for _, joueur := range joueurs {
		joueurTotal := CalculateHand(joueur.hand)
		if (joueurTotal > croupierTotal) && (joueurTotal <= 21) || (croupierTotal > 21) {
			fmt.Printf("Le joueur %s a la main la plus forte!\n", joueur.name)
			fmt.Printf("La valeur de la main du croupier est: %d\n", croupierTotal)
		} else {
			fmt.Printf("Le croupier avait une main plus forte que le joueur %s!\n", joueur.name)
			fmt.Printf("La valeur de la main du croupier est: %d\n", croupierTotal)
		}
	}
}

func intToString(number int) string {
	return strconv.Itoa(number)
}

func CreatePlayers(n int) []Player {
	var joueurs []Player
	for i := 1; i <= n; i++ {
		joueur := NewPlayer(i, "Joueur n°"+intToString(i))
		joueurs = append(joueurs, joueur)
	}
	return joueurs
}

func CreateCroupier() Player {
	croupier := NewPlayer(0, "Croupier")
	return croupier
}
