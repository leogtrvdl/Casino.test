package games

import (
	"fmt"
	"math/rand"
	"time"
)

// creation d'un joueur avec une main vide
func NewPlayer(Number int, Name string) Player {
	return Player{
		number: Number,
		name:   Name,
		hand:   []Card{},
	}
}

// On crée le jeu de carte
func createDeck() []Card {
	symbols := []string{"Coeur", "Trèfle", "Carreau", "Pique"}
	values := []string{"As", "Deux", "Trois", "Quatre", "Cinq", "Six", "Sept", "Huit", "Neuf", "Dix", "Valet", "Dame", "Roi"}

	var deck []Card

	for _, symbol := range symbols {
		for _, value := range values {
			card := Card{value, symbol}
			deck = append(deck, card)
		}
	}

	return deck
}

// On mélange les cartes du jeu
func shuffle(deck []Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
}

// Fonction pour donner une carte aléatoire à un joueur
func giveRandomCard(player *Player, deck *[]Card) {
	if len(*deck) == 0 {
		fmt.Println("Pas assez de cartes dans le jeu.")
		return
	}

	// Choix aléatoire d'une carte dans le jeu
	randomIndex := rand.Intn(len(*deck))
	card := (*deck)[randomIndex]

	// Ajout de la carte à la main du joueur
	player.hand = append(player.hand, card)

	// Retrait de la carte du jeu
	*deck = append((*deck)[:randomIndex], (*deck)[randomIndex+1:]...)
}

// Méthode pour afficher les détails d'un joueur
func (p Player) PrintDetails() {
	fmt.Printf("Joueur %d: %s\n", p.number, p.name)
	fmt.Println("Main:")
	for _, card := range p.hand {
		fmt.Printf(" - %s de %s\n", card.value, card.symbol)
	}
}

// Calcul de la valeur de la main du joueur
func calculateHand(hand []Card) int {
	total := 0

	for _, card := range hand {
		switch card.value {
		case "As":
			total += 1
		case "Deux", "Trois", "Quatre", "Cinq", "Six", "Sept", "Huit", "Neuf", "Dix":
			total += getNameValue(card.value)
		case "Valet", "Dame", "Roi":
			total += 10
		}

	}

	return total
}

// fonction pour passer les valeurs écrites en chiffres
func getNameValue(name string) int {
	values := map[string]int{
		"Deux":   2,
		"Trois":  3,
		"Quatre": 4,
		"Cinq":   5,
		"Six":    6,
		"Sept":   7,
		"Huit":   8,
		"Neuf":   9,
		"Dix":    10,
	}
	return values[name]
}
