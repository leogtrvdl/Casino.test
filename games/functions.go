package games

import (
	"fmt"
	"math/rand"
	"time"
)

// creation d'un joueur avec une main vide
func NewPlayer(Number int, Name string) Player {
	return Player{
		Number: Number,
		Name:   Name,
		Hand:   []Card{},
	}
}

// On crée le jeu de carte
func CreateDeck() []Card {
	symbols := []string{"coeur", "trefle", "carreau", "pique"}
	values := []string{"as", "deux", "trois", "quatre", "cinq", "six", "sept", "huit", "neuf", "dix", "valet", "dame", "roi"}

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
func Shuffle(deck []Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
}

// Fonction pour donner une carte aléatoire à un joueur
func GiveRandomCard(player *Player, deck *[]Card) {
	if len(*deck) == 0 {
		fmt.Println("Pas assez de cartes dans le jeu.")
		return
	}

	// Choix aléatoire d'une carte dans le jeu
	randomIndex := rand.Intn(len(*deck))
	card := (*deck)[randomIndex]

	// Ajout de la carte à la main du joueur
	player.Hand = append(player.Hand, card)

	// Retrait de la carte du jeu
	*deck = append((*deck)[:randomIndex], (*deck)[randomIndex+1:]...)
}

// Méthode pour afficher les détails d'un joueur
func (p Player) PrintDetails() {
	fmt.Printf("Joueur %d: %s\n", p.Number, p.Name)
	fmt.Println("Main:")
	for _, card := range p.Hand {
		fmt.Printf(" - %s de %s\n", card.Value, card.Symbol)
	}
}

// Calcul de la valeur de la main du joueur
func CalculateHand(hand []Card) int {
	total := 0

	for _, card := range hand {
		switch card.Value {
		case "As":
			total += 1
		case "Deux", "Trois", "Quatre", "Cinq", "Six", "Sept", "Huit", "Neuf", "Dix":
			total += getNameValue(card.Value)
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

func (p *Player) GetHand() string {
	hand := fmt.Sprintf("Joueur %d (%s) main:\n", p.Number, p.Name)
	for _, card := range p.Hand {
		hand += fmt.Sprintf("%s de %s\n", card.Value, card.Symbol)
	}
	return hand
}

func (p *Player) GetCroupierHand() string {
	hand := "Croupier main:\n"
	for _, card := range p.Hand {
		hand += fmt.Sprintf("%s de %s\n", card.Value, card.Symbol)
	}
	return hand
}
