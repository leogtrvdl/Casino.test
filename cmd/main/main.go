package main

import (
	"Casinotest/games"
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

type GameState int

const (
	Menu GameState = iota
	Roulette
	Blackjack
	GamesMenu
	APropos
	Joindre
)

type Game struct {
	//images
	img              *ebiten.Image
	ballImg          *ebiten.Image
	rouletteTableImg *ebiten.Image
	background       *ebiten.Image
	dosCartes        *ebiten.Image
	piqueAs          *ebiten.Image
	piqueDeux        *ebiten.Image
	piqueTrois       *ebiten.Image
	piqueQuatre      *ebiten.Image
	piqueCinq        *ebiten.Image
	piqueSix         *ebiten.Image
	piqueSept        *ebiten.Image
	piqueHuit        *ebiten.Image
	piqueNeuf        *ebiten.Image
	piqueDix         *ebiten.Image
	piqueValet       *ebiten.Image
	piqueDame        *ebiten.Image
	piqueRoi         *ebiten.Image
	carreauAs        *ebiten.Image
	carreauDeux      *ebiten.Image
	carreauTrois     *ebiten.Image
	carreauQuatre    *ebiten.Image
	carreauCinq      *ebiten.Image
	carreauSix       *ebiten.Image
	carreauSept      *ebiten.Image
	carreauHuit      *ebiten.Image
	carreauNeuf      *ebiten.Image
	carreauDix       *ebiten.Image
	carreauValet     *ebiten.Image
	carreauDame      *ebiten.Image
	carreauRoi       *ebiten.Image
	coeurAs          *ebiten.Image
	coeurDeux        *ebiten.Image
	coeurTrois       *ebiten.Image
	coeurQuatre      *ebiten.Image
	coeurCinq        *ebiten.Image
	coeurSix         *ebiten.Image
	coeurSept        *ebiten.Image
	coeurHuit        *ebiten.Image
	coeurNeuf        *ebiten.Image
	coeurDix         *ebiten.Image
	coeurValet       *ebiten.Image
	coeurDame        *ebiten.Image
	coeurRoi         *ebiten.Image
	trefleAs         *ebiten.Image
	trefleDeux       *ebiten.Image
	trefleTrois      *ebiten.Image
	trefleQuatre     *ebiten.Image
	trefleCinq       *ebiten.Image
	trefleSix        *ebiten.Image
	trefleSept       *ebiten.Image
	trefleHuit       *ebiten.Image
	trefleNeuf       *ebiten.Image
	trefleDix        *ebiten.Image
	trefleValet      *ebiten.Image
	trefleDame       *ebiten.Image
	trefleRoi        *ebiten.Image

	//pas images
	angle            float64
	ballAngle        float64
	ballRadius       float64
	state            GameState
	rouletteStart    time.Time
	isSpinning       bool
	ballVisible      bool
	ballX            float64
	ballY            float64
	isTimerStarted   bool
	roulValue        int
	roulColor        string
	roulTier         int
	roulLigne        int
	roulDemi         int
	blackjackStart   bool
	isPlayerSelected bool
	isPlayersTurn    bool
	playerNumber     int
	symbolBj         string
	valueBj          string
	croupier         games.Player
	players          []games.Player
	croupierHand     string
	playersHand      []string
	cardToShow       string
	cardsToShow      []string
	currentPlayer    int
	cards            []games.Card
	handValue        int
}

func (g *Game) Update() error {
	if g.state == Menu {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			switch {
			case x >= 20 && x <= 120 && y >= 70 && y < 110:
				g.state = GamesMenu
			case x >= 20 && x <= 120 && y >= 110 && y < 145:
				g.state = APropos
			case x >= 20 && x <= 120 && y >= 145 && y < 175:
				g.state = Joindre
			}
		}
	} else if g.state == GamesMenu {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			switch {
			case x >= 20 && x <= 120 && y >= 95 && y < 105:
				g.state = Menu
			case x >= 20 && x <= 120 && y >= 165 && y < 175:
				g.state = Roulette
			case x >= 20 && x <= 120 && y >= 200 && y < 210:
				g.state = Blackjack
			}
		}
	} else if g.state == APropos {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			switch {
			case x >= 20 && x <= 120 && y >= 70 && y < 105:
				g.state = Menu
			}
		}
	} else if g.state == Joindre {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			switch {
			case x >= 20 && x <= 120 && y >= 60 && y < 75:
				g.state = Menu
			}
		}
	} else if g.state == Roulette {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			switch {
			case x >= 10 && x <= 160 && y >= 140 && y <= 175:
				g.isSpinning = true
				g.ballVisible = true
				g.rouletteStart = time.Now()
				g.isTimerStarted = true
				g.ballAngle = rand.Float64() * 2 * math.Pi // angle de départ rdm
				g.ballRadius = 90.0                        // Radius balle
			case x >= 10 && x <= 160 && y >= 75 && y <= 115:
				g.state = GamesMenu
			}
		}
		if g.isSpinning {
			g.angle += 0.05
			g.ballAngle += 0.1

			if time.Since(g.rouletteStart).Seconds() > 10 {
				g.isSpinning = false
				g.ballVisible = false
				g.angle = 0
				numero := games.PlayRoulette()
				g.roulValue, g.roulColor, g.roulTier, g.roulLigne, g.roulDemi = numero.Value, numero.Color, numero.Tier, numero.Ligne, numero.Demi
				//tp la balle au num indiqué (faut raccourcire)

				switch numero.Value {
				case 0:
					g.ballX = 490
					g.ballY = 80
				case 1:
					g.ballX = 140
					g.ballY = 150
				case 2:
					g.ballX = 528
					g.ballY = 175
				case 3:
					g.ballX = 465
					g.ballY = 62
				case 4:
					g.ballX = 525
					g.ballY = 140
				case 5:
					g.ballX = 350
					g.ballY = 225
				case 6:
					g.ballX = 490
					g.ballY = 230
				case 7:
					g.ballX = 395
					g.ballY = 58
				case 8:
					g.ballX = 395
					g.ballY = 255
				case 9:
					g.ballX = 340
					g.ballY = 100
				case 10:
					g.ballX = 359
					g.ballY = 235
				case 11:
					g.ballX = 427
					g.ballY = 260
				case 12:
					g.ballX = 430
					g.ballY = 55
				case 13:
					g.ballX = 460
					g.ballY = 250
				case 14:
					g.ballX = 325
					g.ballY = 130
				case 15:
					g.ballX = 515
					g.ballY = 105
				case 16:
					g.ballX = 330
					g.ballY = 200
				case 17:
					g.ballX = 515
					g.ballY = 210
				case 18:
					g.ballX = 362
					g.ballY = 75
				case 19:
					g.ballX = 522
					g.ballY = 122
				case 20:
					g.ballX = 321
					g.ballY = 147
				case 21:
					g.ballX = 530
					g.ballY = 160
				case 22:
					g.ballX = 350
					g.ballY = 85
				case 23:
					g.ballX = 377
					g.ballY = 247
				case 24:
					g.ballX = 340
					g.ballY = 215
				case 25:
					g.ballX = 520
					g.ballY = 190
				case 26:
					g.ballX = 480
					g.ballY = 70
				case 27:
					g.ballX = 505
					g.ballY = 220
				case 28:
					g.ballX = 410
					g.ballY = 55
				case 29:
					g.ballX = 380
					g.ballY = 65
				case 30:
					g.ballX = 410
					g.ballY = 260
				case 31:
					g.ballX = 330
					g.ballY = 110
				case 32:
					g.ballX = 505
					g.ballY = 90
				case 33:
					g.ballX = 322
					g.ballY = 180
				case 34:
					g.ballX = 477
					g.ballY = 243
				case 35:
					g.ballX = 445
					g.ballY = 55
				case 36:
					g.ballX = 447
					g.ballY = 260
				}
			}
		}
	} else if g.state == Blackjack {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			switch {
			case x >= 20 && x <= 120 && y >= 60 && y < 75:
				g.state = GamesMenu
				g.blackjackStart = false
				g.isPlayerSelected = false
			case x >= 10 && x <= 160 && y >= 165 && y <= 185:
				g.blackjackStart = true
			case g.blackjackStart && x >= 200 && x <= 300 && y >= 155 && y <= 175:
				g.playerNumber = 1
				g.isPlayerSelected = true
			case g.blackjackStart && x >= 200 && x <= 300 && y >= 185 && y <= 205:
				g.playerNumber = 2
				g.isPlayerSelected = true
			case g.blackjackStart && x >= 200 && x <= 300 && y >= 215 && y <= 235:
				g.playerNumber = 3
				g.isPlayerSelected = true
			case g.blackjackStart && x >= 200 && x <= 300 && y >= 245 && y <= 265:
				g.playerNumber = 4
				g.isPlayerSelected = true
			case g.blackjackStart && x >= 200 && x <= 300 && y >= 275 && y <= 295:
				g.playerNumber = 5
				g.isPlayerSelected = true
			case x >= 20 && x <= 120 && y >= 240 && y < 255:
				games.GiveRandomCard(&g.players[g.currentPlayer], &g.cards)
				g.playersHand[g.currentPlayer] = g.players[g.currentPlayer].GetHand()
				g.handValue = g.players[g.currentPlayer].CalculateHandValue()
			case x >= 160 && x <= 260 && y >= 240 && y < 255:
				if g.currentPlayer < g.playerNumber {
					g.currentPlayer++
				} else if g.currentPlayer >= g.playerNumber {
					g.currentPlayer = 0
				}
				g.handValue = g.players[g.currentPlayer].CalculateHandValue()

			}
			if g.isPlayerSelected && g.blackjackStart && !g.isPlayersTurn {
				//creation du deck de jeu
				g.cards = games.CreateDeck()
				games.Shuffle(g.cards)

				//creation des joueurs
				g.croupier = games.CreateCroupier()
				g.players = games.CreatePlayers(g.playerNumber)
				g.currentPlayer = 0

				//attribution de cartes aux joueurs
				for j := 0; j < 2; j++ {
					games.GiveRandomCard(&g.croupier, &g.cards)
				}
				for i := range g.players {
					games.GiveRandomCard(&g.players[i], &g.cards)
				}

				//Indiquer au jeu que les cartes ont été données
				g.isPlayersTurn = true

				g.currentPlayer = 0

				// Obtenir la main du croupier sous forme de chaîne
				g.croupierHand = g.croupier.GetCroupierHand()

				// Obtenir les mains des joueurs sous forme de chaîne
				for _, player := range g.players {
					g.playersHand = append(g.playersHand, player.GetHand())
				}

			}
		}
	} else {
		g.angle += 0.01
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	screen.Fill(color.White)

	switch {
	case g.state == Menu:
		g.drawMenu(screen)
	case g.state == Roulette:
		g.drawRoulette(screen)
	case g.state == Blackjack:
		g.drawBlackjack(screen)
	case g.state == GamesMenu:
		g.drawGamesMenu(screen)
	case g.state == Joindre:
		g.drawJoindre(screen)
	case g.state == APropos:
		g.drawAPropos(screen)
	}
}

func (g *Game) drawMenu(screen *ebiten.Image) {
	screen.DrawImage(g.background, nil)
	g.drawOption(screen, "Nos jeux", 20, 105, true, false)
	g.drawOption(screen, "A propos", 20, 140, true, false)
	g.drawOption(screen, "Me joindre", 20, 175, true, false)
}

func (g *Game) drawAPropos(screen *ebiten.Image) {
	screen.DrawImage(g.background, nil)
	g.drawOption(screen, "Retour au menu", 20, 105, true, false)
	g.drawOption(screen, "Ce casino est un projet personnel \nrealiser dans le cadre de \nma formation afin de developper ma \nconnaissance du language Go", 180, 135, false, false)
}

func (g *Game) drawJoindre(screen *ebiten.Image) {
	screen.DrawImage(g.background, nil)
	g.drawOption(screen, "Retour au menu", 20, 70, true, false)
	g.drawOption(screen, "Pour me joindre : ", 20, 140, false, false)
	g.drawOption(screen, "Mon Linkedin  -> ", 60, 175, false, false)
	g.drawOption(screen, "Mon Email     -> ", 60, 210, false, false)
	g.drawOption(screen, "Mon Github    ->  ", 60, 245, false, false)
	g.drawOption(screen, "https://www.linkedin.com/in/leo-gautier-vidal-850702239/", 200, 175, true, true)
	g.drawOption(screen, "leogautiervidal@gmail.com", 200, 210, true, true)
	g.drawOption(screen, "https://github.com/leogtrvdl", 200, 245, true, true)
}

func (g *Game) drawGamesMenu(screen *ebiten.Image) {
	screen.DrawImage(g.background, nil)
	g.drawOption(screen, "Retour au menu", 20, 105, true, false)
	g.drawOption(screen, "Roulette", 20, 175, true, false)
	g.drawOption(screen, "Blackjack", 20, 210, true, false)
}

func (g *Game) drawBlackjack(screen *ebiten.Image) {
	screen.DrawImage(g.background, nil)
	g.drawOption(screen, "Retour aux jeux", 20, 70, true, false)
	if !g.blackjackStart {
		g.drawOption(screen, "Rejoindre la partie", 20, 175, true, false)
	}
	if g.blackjackStart && !g.isPlayerSelected {
		g.drawOption(screen, "Combien de joueurs pour cette partie? ", 200, 125, false, false)
		g.drawOption(screen, "1. Player ", 200, 160, true, false)
		g.drawOption(screen, "2. Player ", 200, 190, true, false)
		g.drawOption(screen, "3. Player ", 200, 220, true, false)
		g.drawOption(screen, "4. Player ", 200, 250, true, false)
		g.drawOption(screen, "5. Player ", 200, 280, true, false)
	}
	if g.blackjackStart && g.isPlayerSelected {

		g.drawOption(screen, "Tirer une carte", 20, 250, true, false)
		g.drawOption(screen, "Passer", 160, 250, true, false)
		g.drawOption(screen, "tour du joueur ", 40, 200, false, false)
		g.drawOption(screen, intToString(g.currentPlayer+1), 180, 200, false, false)
		g.drawOption(screen, intToString(g.handValue), 260, 200, false, false)
	}

	if g.isPlayersTurn && g.blackjackStart && !(g.currentPlayer >= g.playerNumber) {
		cards, err := games.ExtractCardValues(g.croupierHand)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		xPosition := 250.0 // Position de départ pour l'affichage des cartes
		yPosition := 20.0
		spacingX := 25.0 // Espacement entre les cartes
		spacingY := 10.0
		scale := 0.8

		for _, card := range cards {
			cardName := fmt.Sprintf("%s%s", card.Symbol, games.CapitalizeFirstLetter(card.Value))
			g.drawCard(screen, cardName, xPosition, yPosition, scale)
			xPosition += spacingX // Ajustez l'espace entre les cartes si nécessaire
			yPosition += spacingY
		}

		// Afficher la valeur de la main du croupier
		croupierHandValue := g.croupier.CalculateHandValue()
		text.Draw(screen, fmt.Sprintf("%d", croupierHandValue), text.FaceWithLineHeight(basicfont.Face7x13, 15), int(xPosition)+100, int(yPosition), color.RGBA{255, 255, 255, 255})

		switch g.currentPlayer {
		case 0:
			playerCards, err := games.ExtractCardValues(g.playersHand[0])
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			xPosition := 20.0 // Position de départ pour l'affichage des cartes
			yPosition := 300.0
			spacingX := 25.0 // Espacement entre les cartes
			spacingY := 10.0
			scale := 0.8

			for _, card := range playerCards {
				cardName := fmt.Sprintf("%s%s", card.Symbol, games.CapitalizeFirstLetter(card.Value))
				g.drawCard(screen, cardName, xPosition, yPosition, scale)
				xPosition += spacingX // Ajustez l'espace entre les cartes si nécessaire
				yPosition += spacingY
			}
		case 1:
			playerCards, err := games.ExtractCardValues(g.playersHand[1])
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			xPosition := 20.0 // Position de départ pour l'affichage des cartes
			yPosition := 300.0
			spacingX := 25.0 // Espacement entre les cartes
			spacingY := 10.0
			scale := 0.8

			for _, card := range playerCards {
				cardName := fmt.Sprintf("%s%s", card.Symbol, games.CapitalizeFirstLetter(card.Value))
				g.drawCard(screen, cardName, xPosition, yPosition, scale)
				xPosition += spacingX // Ajustez l'espace entre les cartes si nécessaire
				yPosition += spacingY
			}
		case 2:
			playerCards, err := games.ExtractCardValues(g.playersHand[2])
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			xPosition := 20.0 // Position de départ pour l'affichage des cartes
			yPosition := 300.0
			spacingX := 25.0 // Espacement entre les cartes
			spacingY := 10.0
			scale := 0.8

			for _, card := range playerCards {
				cardName := fmt.Sprintf("%s%s", card.Symbol, games.CapitalizeFirstLetter(card.Value))
				g.drawCard(screen, cardName, xPosition, yPosition, scale)
				xPosition += spacingX // Ajustez l'espace entre les cartes si nécessaire
				yPosition += spacingY
			}
		case 3:
			playerCards, err := games.ExtractCardValues(g.playersHand[3])
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			xPosition := 20.0 // Position de départ pour l'affichage des cartes
			yPosition := 300.0
			spacingX := 25.0 // Espacement entre les cartes
			spacingY := 10.0
			scale := 0.8

			for _, card := range playerCards {
				cardName := fmt.Sprintf("%s%s", card.Symbol, games.CapitalizeFirstLetter(card.Value))
				g.drawCard(screen, cardName, xPosition, yPosition, scale)
				xPosition += spacingX // Ajustez l'espace entre les cartes si nécessaire
				yPosition += spacingY
			}
		case 4:
			playerCards, err := games.ExtractCardValues(g.playersHand[4])
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			xPosition := 20.0 // Position de départ pour l'affichage des cartes
			yPosition := 300.0
			spacingX := 25.0 // Espacement entre les cartes
			spacingY := 10.0
			scale := 0.8

			for _, card := range playerCards {
				cardName := fmt.Sprintf("%s%s", card.Symbol, games.CapitalizeFirstLetter(card.Value))
				g.drawCard(screen, cardName, xPosition, yPosition, scale)
				xPosition += spacingX // Ajustez l'espace entre les cartes si nécessaire
				yPosition += spacingY
			}
		}
	}
}

func (g *Game) drawCard(screen *ebiten.Image, cardName string, x, y, scale float64) {
	cartePos := &ebiten.DrawImageOptions{}
	cartePos.GeoM.Scale(scale, scale) // Mise à l'échelle de la carte
	cartePos.GeoM.Translate(x, y)

	switch cardName {
	// Pique
	case "piqueAs":
		screen.DrawImage(g.piqueAs, cartePos)
	case "piqueDeux":
		screen.DrawImage(g.piqueDeux, cartePos)
	case "piqueTrois":
		screen.DrawImage(g.piqueTrois, cartePos)
	case "piqueQuatre":
		screen.DrawImage(g.piqueQuatre, cartePos)
	case "piqueCinq":
		screen.DrawImage(g.piqueCinq, cartePos)
	case "piqueSix":
		screen.DrawImage(g.piqueSix, cartePos)
	case "piqueSept":
		screen.DrawImage(g.piqueSept, cartePos)
	case "piqueHuit":
		screen.DrawImage(g.piqueHuit, cartePos)
	case "piqueNeuf":
		screen.DrawImage(g.piqueNeuf, cartePos)
	case "piqueDix":
		screen.DrawImage(g.piqueDix, cartePos)
	case "piqueValet":
		screen.DrawImage(g.piqueValet, cartePos)
	case "piqueDame":
		screen.DrawImage(g.piqueDame, cartePos)
	case "piqueRoi":
		screen.DrawImage(g.piqueRoi, cartePos)
	// Carreau
	case "carreauAs":
		screen.DrawImage(g.carreauAs, cartePos)
	case "carreauDeux":
		screen.DrawImage(g.carreauDeux, cartePos)
	case "carreauTrois":
		screen.DrawImage(g.carreauTrois, cartePos)
	case "carreauQuatre":
		screen.DrawImage(g.carreauQuatre, cartePos)
	case "carreauCinq":
		screen.DrawImage(g.carreauCinq, cartePos)
	case "carreauSix":
		screen.DrawImage(g.carreauSix, cartePos)
	case "carreauSept":
		screen.DrawImage(g.carreauSept, cartePos)
	case "carreauHuit":
		screen.DrawImage(g.carreauHuit, cartePos)
	case "carreauNeuf":
		screen.DrawImage(g.carreauNeuf, cartePos)
	case "carreauDix":
		screen.DrawImage(g.carreauDix, cartePos)
	case "carreauValet":
		screen.DrawImage(g.carreauValet, cartePos)
	case "carreauDame":
		screen.DrawImage(g.carreauDame, cartePos)
	case "carreauRoi":
		screen.DrawImage(g.carreauRoi, cartePos)
	// Coeur
	case "coeurAs":
		screen.DrawImage(g.coeurAs, cartePos)
	case "coeurDeux":
		screen.DrawImage(g.coeurDeux, cartePos)
	case "coeurTrois":
		screen.DrawImage(g.coeurTrois, cartePos)
	case "coeurQuatre":
		screen.DrawImage(g.coeurQuatre, cartePos)
	case "coeurCinq":
		screen.DrawImage(g.coeurCinq, cartePos)
	case "coeurSix":
		screen.DrawImage(g.coeurSix, cartePos)
	case "coeurSept":
		screen.DrawImage(g.coeurSept, cartePos)
	case "coeurHuit":
		screen.DrawImage(g.coeurHuit, cartePos)
	case "coeurNeuf":
		screen.DrawImage(g.coeurNeuf, cartePos)
	case "coeurDix":
		screen.DrawImage(g.coeurDix, cartePos)
	case "coeurValet":
		screen.DrawImage(g.coeurValet, cartePos)
	case "coeurDame":
		screen.DrawImage(g.coeurDame, cartePos)
	case "coeurRoi":
		screen.DrawImage(g.coeurRoi, cartePos)
	// Trèfle
	case "trefleAs":
		screen.DrawImage(g.trefleAs, cartePos)
	case "trefleDeux":
		screen.DrawImage(g.trefleDeux, cartePos)
	case "trefleTrois":
		screen.DrawImage(g.trefleTrois, cartePos)
	case "trefleQuatre":
		screen.DrawImage(g.trefleQuatre, cartePos)
	case "trefleCinq":
		screen.DrawImage(g.trefleCinq, cartePos)
	case "trefleSix":
		screen.DrawImage(g.trefleSix, cartePos)
	case "trefleSept":
		screen.DrawImage(g.trefleSept, cartePos)
	case "trefleHuit":
		screen.DrawImage(g.trefleHuit, cartePos)
	case "trefleNeuf":
		screen.DrawImage(g.trefleNeuf, cartePos)
	case "trefleDix":
		screen.DrawImage(g.trefleDix, cartePos)
	case "trefleValet":
		screen.DrawImage(g.trefleValet, cartePos)
	case "trefleDame":
		screen.DrawImage(g.trefleDame, cartePos)
	case "trefleRoi":
		screen.DrawImage(g.trefleRoi, cartePos)
	}
}

func (g *Game) drawRoulette(screen *ebiten.Image) {
	w, h := g.img.Size()
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(-float64(w)/4, -float64(h)/4)
	if g.isSpinning {
		op.GeoM.Rotate(g.angle)
	}

	additionalTranslationX := float64(screen.Bounds().Dx()) / 6
	additionalTranslationY := float64(screen.Bounds().Dy()) / 6
	op.GeoM.Translate(float64(screen.Bounds().Dx())/2+additionalTranslationX, float64(screen.Bounds().Dy())/2-additionalTranslationY)

	screen.DrawImage(g.background, nil)
	screen.DrawImage(g.img, op)

	// Draw ball
	if g.ballVisible {
		ballOp := &ebiten.DrawImageOptions{}
		ballW, ballH := g.ballImg.Size()
		ballOp.GeoM.Translate(-float64(ballW)/2, -float64(ballH)/2)

		// Calculate ball position
		ballX := float64(screen.Bounds().Dx())/2 + additionalTranslationX + g.ballRadius*math.Cos(g.ballAngle)
		ballY := float64(screen.Bounds().Dy())/2 - additionalTranslationY + g.ballRadius*math.Sin(g.ballAngle)

		ballOp.GeoM.Translate(ballX, ballY)
		screen.DrawImage(g.ballImg, ballOp)
	}

	// Draw gambling table
	ballOp := &ebiten.DrawImageOptions{}
	ballOp.GeoM.Translate(175, 270)
	screen.DrawImage(g.rouletteTableImg, ballOp)

	// Draw options with hover effects
	g.drawOption(screen, "Lancer une balle", 20, 175, true, false)
	g.drawOption(screen, "Retour aux jeux", 20, 105, true, false)

	// Création des message a montrer + passage d'int en string
	msgVa := "Le nombre : " + intToString(g.roulValue)
	msgCo := "La couleur : " + g.roulColor
	msgTi := "Le tier : " + intToString(g.roulTier)
	msgLi := "La ligne : " + intToString(g.roulLigne)
	msgDe := "La moitie de table : " + intToString(g.roulDemi)

	// Text des gagnants de la roulette
	if time.Since(g.rouletteStart).Seconds() > 10 && g.isTimerStarted {
		text.Draw(screen, "Les possibilites gagnantes sont : ", text.FaceWithLineHeight(basicfont.Face7x13, 15), 20, 320, color.RGBA{255, 255, 255, 255})
		text.Draw(screen, msgVa, text.FaceWithLineHeight(basicfont.Face7x13, 15), 20, 360, color.RGBA{255, 255, 255, 255})
		text.Draw(screen, msgCo, text.FaceWithLineHeight(basicfont.Face7x13, 15), 20, 380, color.RGBA{255, 255, 255, 255})
		text.Draw(screen, msgTi, text.FaceWithLineHeight(basicfont.Face7x13, 15), 20, 400, color.RGBA{255, 255, 255, 255})
		text.Draw(screen, msgLi, text.FaceWithLineHeight(basicfont.Face7x13, 15), 20, 420, color.RGBA{255, 255, 255, 255})
		text.Draw(screen, msgDe, text.FaceWithLineHeight(basicfont.Face7x13, 15), 20, 440, color.RGBA{255, 255, 255, 255})
		//ball after
		ballOp := &ebiten.DrawImageOptions{}
		ballW, ballH := g.ballImg.Size()
		ballOp.GeoM.Translate(-float64(ballW)/2, -float64(ballH)/2)

		ballOp.GeoM.Translate(g.ballX, g.ballY)
		screen.DrawImage(g.ballImg, ballOp)
	}
}

func (g *Game) drawOption(screen *ebiten.Image, msg string, x, y int, shouldHighlightOnHover, isHyperlink bool) {
	regularFace := basicfont.Face7x13
	boldFace := text.FaceWithLineHeight(basicfont.Face7x13, 15) // larger font for hover effect

	// Check if cursor is over the text
	isHovered := isCursorOverText(x, y, len(msg)*7, 13)

	// Draw text with or without hover effect
	if isHovered && shouldHighlightOnHover {
		text.Draw(screen, msg, boldFace, x, y, color.RGBA{255, 255, 255, 255})                                        // Highlighted color
		ebitenutil.DrawRect(screen, float64(x), float64(y+5), float64(len(msg)*8), 1, color.RGBA{255, 255, 255, 255}) // Underline
	} else {
		text.Draw(screen, msg, regularFace, x, y, color.RGBA{255, 255, 255, 255}) // Regular color
	}

	// Open URL if text is clicked and it is a hyperlink
	if isHovered && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && isHyperlink {
		openURL(msg)
	}
}

func openURL(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	}
	if err != nil {
		panic(err)
	}
}

func isCursorOverText(x, y, textWidth, textHeight int) bool {
	cursorX, cursorY := ebiten.CursorPosition()
	return cursorX >= x && cursorX <= x+textWidth && cursorY >= y-textHeight && cursorY <= y
}

func intToString(number int) string {
	return strconv.Itoa(number)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func setImg() {
	img, _, err := ebitenutil.NewImageFromFile("assets/roulette.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	ballImg, _, err := ebitenutil.NewImageFromFile("assets/balle.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image de la balle : %v", err)
		os.Exit(1)
	}

	rouletteTableImg, _, err := ebitenutil.NewImageFromFile("assets/rouletteTable.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image de la table de la roulette : %v", err)
		os.Exit(1)
	}

	background, _, err := ebitenutil.NewImageFromFile("assets/background.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image de la table de la roulette : %v", err)
		os.Exit(1)
	}

	dosCartes, _, err := ebitenutil.NewImageFromFile("assets/cartes/dos_cartes.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	trefleAs, _, err := ebitenutil.NewImageFromFile("assets/cartes/trefle/as.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	trefleDeux, _, err := ebitenutil.NewImageFromFile("assets/cartes/trefle/deux.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	trefleTrois, _, err := ebitenutil.NewImageFromFile("assets/cartes/trefle/trois.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	trefleQuatre, _, err := ebitenutil.NewImageFromFile("assets/cartes/trefle/quatre.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	trefleCinq, _, err := ebitenutil.NewImageFromFile("assets/cartes/trefle/cinq.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	trefleSix, _, err := ebitenutil.NewImageFromFile("assets/cartes/trefle/six.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	trefleSept, _, err := ebitenutil.NewImageFromFile("assets/cartes/trefle/sept.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	trefleHuit, _, err := ebitenutil.NewImageFromFile("assets/cartes/trefle/huit.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	trefleNeuf, _, err := ebitenutil.NewImageFromFile("assets/cartes/trefle/neuf.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	trefleDix, _, err := ebitenutil.NewImageFromFile("assets/cartes/trefle/dix.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	trefleValet, _, err := ebitenutil.NewImageFromFile("assets/cartes/trefle/valet.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	trefleDame, _, err := ebitenutil.NewImageFromFile("assets/cartes/trefle/dame.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	trefleRoi, _, err := ebitenutil.NewImageFromFile("assets/cartes/trefle/roi.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	piqueAs, _, err := ebitenutil.NewImageFromFile("assets/cartes/pique/as.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	piqueDeux, _, err := ebitenutil.NewImageFromFile("assets/cartes/pique/deux.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	piqueTrois, _, err := ebitenutil.NewImageFromFile("assets/cartes/pique/trois.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	piqueQuatre, _, err := ebitenutil.NewImageFromFile("assets/cartes/pique/quatre.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	piqueCinq, _, err := ebitenutil.NewImageFromFile("assets/cartes/pique/cinq.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	piqueSix, _, err := ebitenutil.NewImageFromFile("assets/cartes/pique/six.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	piqueSept, _, err := ebitenutil.NewImageFromFile("assets/cartes/pique/sept.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	piqueHuit, _, err := ebitenutil.NewImageFromFile("assets/cartes/pique/huit.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	piqueNeuf, _, err := ebitenutil.NewImageFromFile("assets/cartes/pique/neuf.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	piqueDix, _, err := ebitenutil.NewImageFromFile("assets/cartes/pique/dix.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	piqueValet, _, err := ebitenutil.NewImageFromFile("assets/cartes/pique/valet.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	piqueDame, _, err := ebitenutil.NewImageFromFile("assets/cartes/pique/dame.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	piqueRoi, _, err := ebitenutil.NewImageFromFile("assets/cartes/pique/roi.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	carreauAs, _, err := ebitenutil.NewImageFromFile("assets/cartes/carreau/as.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	carreauDeux, _, err := ebitenutil.NewImageFromFile("assets/cartes/carreau/deux.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	carreauTrois, _, err := ebitenutil.NewImageFromFile("assets/cartes/carreau/trois.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	carreauQuatre, _, err := ebitenutil.NewImageFromFile("assets/cartes/carreau/quatre.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	carreauCinq, _, err := ebitenutil.NewImageFromFile("assets/cartes/carreau/cinq.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	carreauSix, _, err := ebitenutil.NewImageFromFile("assets/cartes/carreau/six.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	carreauSept, _, err := ebitenutil.NewImageFromFile("assets/cartes/carreau/sept.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	carreauHuit, _, err := ebitenutil.NewImageFromFile("assets/cartes/carreau/huit.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	carreauNeuf, _, err := ebitenutil.NewImageFromFile("assets/cartes/carreau/neuf.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	carreauDix, _, err := ebitenutil.NewImageFromFile("assets/cartes/carreau/dix.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	carreauValet, _, err := ebitenutil.NewImageFromFile("assets/cartes/carreau/valet.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	carreauDame, _, err := ebitenutil.NewImageFromFile("assets/cartes/carreau/dame.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	carreauRoi, _, err := ebitenutil.NewImageFromFile("assets/cartes/carreau/roi.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	coeurAs, _, err := ebitenutil.NewImageFromFile("assets/cartes/coeur/as.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	coeurDeux, _, err := ebitenutil.NewImageFromFile("assets/cartes/coeur/deux.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	coeurTrois, _, err := ebitenutil.NewImageFromFile("assets/cartes/coeur/trois.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	coeurQuatre, _, err := ebitenutil.NewImageFromFile("assets/cartes/coeur/quatre.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	coeurCinq, _, err := ebitenutil.NewImageFromFile("assets/cartes/coeur/cinq.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	coeurSix, _, err := ebitenutil.NewImageFromFile("assets/cartes/coeur/six.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	coeurSept, _, err := ebitenutil.NewImageFromFile("assets/cartes/coeur/sept.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	coeurHuit, _, err := ebitenutil.NewImageFromFile("assets/cartes/coeur/huit.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	coeurNeuf, _, err := ebitenutil.NewImageFromFile("assets/cartes/coeur/neuf.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	coeurDix, _, err := ebitenutil.NewImageFromFile("assets/cartes/coeur/dix.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	coeurValet, _, err := ebitenutil.NewImageFromFile("assets/cartes/coeur/valet.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	coeurDame, _, err := ebitenutil.NewImageFromFile("assets/cartes/coeur/dame.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	coeurRoi, _, err := ebitenutil.NewImageFromFile("assets/cartes/coeur/roi.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	CardImage := make(map[string]*ebiten.Image)

	addCardImage := func(path, name string) {
		img, _, err := ebitenutil.NewImageFromFile(path)
		if err != nil {
			log.Fatalf("Erreur lors du chargement de l'image %s : %v", name, err)
			os.Exit(1)
		}
		CardImage[name] = img
	}

	//trefles
	addCardImage("assets/cartes/trefle/as.png", "trefle_as")
	addCardImage("assets/cartes/trefle/deux.png", "trefle_deux")
	addCardImage("assets/cartes/trefle/trois.png", "trefle_trois")
	addCardImage("assets/cartes/trefle/quatre.png", "trefle_quatre")
	addCardImage("assets/cartes/trefle/cinq.png", "trefle_cinq")
	addCardImage("assets/cartes/trefle/six.png", "trefle_six")
	addCardImage("assets/cartes/trefle/sept.png", "trefle_sept")
	addCardImage("assets/cartes/trefle/huit.png", "trefle_huit")
	addCardImage("assets/cartes/trefle/neuf.png", "trefle_neuf")
	addCardImage("assets/cartes/trefle/dix.png", "trefle_dix")
	addCardImage("assets/cartes/trefle/valet.png", "trefle_valet")
	addCardImage("assets/cartes/trefle/dame.png", "trefle_dame")
	addCardImage("assets/cartes/trefle/roi.png", "trefle_roi")

	//pique
	addCardImage("assets/cartes/pique/as.png", "pique_as")
	addCardImage("assets/cartes/pique/deux.png", "pique_deux")
	addCardImage("assets/cartes/pique/trois.png", "pique_trois")
	addCardImage("assets/cartes/pique/quatre.png", "pique_quatre")
	addCardImage("assets/cartes/pique/cinq.png", "pique_cinq")
	addCardImage("assets/cartes/pique/six.png", "pique_six")
	addCardImage("assets/cartes/pique/sept.png", "pique_sept")
	addCardImage("assets/cartes/pique/huit.png", "pique_huit")
	addCardImage("assets/cartes/pique/neuf.png", "pique_neuf")
	addCardImage("assets/cartes/pique/dix.png", "pique_dix")
	addCardImage("assets/cartes/pique/valet.png", "pique_valet")
	addCardImage("assets/cartes/pique/dame.png", "pique_dame")
	addCardImage("assets/cartes/pique/roi.png", "pique_roi")

	//carreau
	addCardImage("assets/cartes/carreau/as.png", "carreau_as")
	addCardImage("assets/cartes/carreau/deux.png", "carreau_deux")
	addCardImage("assets/cartes/carreau/trois.png", "carreau_trois")
	addCardImage("assets/cartes/carreau/quatre.png", "carreau_quatre")
	addCardImage("assets/cartes/carreau/cinq.png", "carreau_cinq")
	addCardImage("assets/cartes/carreau/six.png", "carreau_six")
	addCardImage("assets/cartes/carreau/sept.png", "carreau_sept")
	addCardImage("assets/cartes/carreau/huit.png", "carreau_huit")
	addCardImage("assets/cartes/carreau/neuf.png", "carreau_neuf")
	addCardImage("assets/cartes/carreau/dix.png", "carreau_dix")
	addCardImage("assets/cartes/carreau/valet.png", "carreau_valet")
	addCardImage("assets/cartes/carreau/dame.png", "carreau_dame")
	addCardImage("assets/cartes/carreau/roi.png", "carreau_roi")

	//coeur
	addCardImage("assets/cartes/coeur/as.png", "coeur_as")
	addCardImage("assets/cartes/coeur/deux.png", "coeur_deux")
	addCardImage("assets/cartes/coeur/trois.png", "coeur_trois")
	addCardImage("assets/cartes/coeur/quatre.png", "coeur_quatre")
	addCardImage("assets/cartes/coeur/cinq.png", "coeur_cinq")
	addCardImage("assets/cartes/coeur/six.png", "coeur_six")
	addCardImage("assets/cartes/coeur/sept.png", "coeur_sept")
	addCardImage("assets/cartes/coeur/huit.png", "coeur_huit")
	addCardImage("assets/cartes/coeur/neuf.png", "coeur_neuf")
	addCardImage("assets/cartes/coeur/dix.png", "coeur_dix")
	addCardImage("assets/cartes/coeur/valet.png", "coeur_valet")
	addCardImage("assets/cartes/coeur/dame.png", "coeur_dame")
	addCardImage("assets/cartes/coeur/roi.png", "coeur_roi")

	game := &Game{
		img:              img,
		ballImg:          ballImg,
		rouletteTableImg: rouletteTableImg,
		state:            Menu,
		background:       background,
		dosCartes:        dosCartes,
		trefleAs:         trefleAs,
		trefleDeux:       trefleDeux,
		trefleTrois:      trefleTrois,
		trefleQuatre:     trefleQuatre,
		trefleCinq:       trefleCinq,
		trefleSix:        trefleSix,
		trefleSept:       trefleSept,
		trefleHuit:       trefleHuit,
		trefleNeuf:       trefleNeuf,
		trefleDix:        trefleDix,
		trefleValet:      trefleValet,
		trefleDame:       trefleDame,
		trefleRoi:        trefleRoi,
		piqueAs:          piqueAs,
		piqueDeux:        piqueDeux,
		piqueTrois:       piqueTrois,
		piqueQuatre:      piqueQuatre,
		piqueCinq:        piqueCinq,
		piqueSix:         piqueSix,
		piqueSept:        piqueSept,
		piqueHuit:        piqueHuit,
		piqueNeuf:        piqueNeuf,
		piqueDix:         piqueDix,
		piqueValet:       piqueValet,
		piqueDame:        piqueDame,
		piqueRoi:         piqueRoi,
		carreauAs:        carreauAs,
		carreauDeux:      carreauDeux,
		carreauTrois:     carreauTrois,
		carreauQuatre:    carreauQuatre,
		carreauCinq:      carreauCinq,
		carreauSix:       carreauSix,
		carreauSept:      carreauSept,
		carreauHuit:      carreauHuit,
		carreauNeuf:      carreauNeuf,
		carreauDix:       carreauDix,
		carreauValet:     carreauValet,
		carreauDame:      carreauDame,
		carreauRoi:       carreauRoi,
		coeurAs:          coeurAs,
		coeurDeux:        coeurDeux,
		coeurTrois:       coeurTrois,
		coeurQuatre:      coeurQuatre,
		coeurCinq:        coeurCinq,
		coeurSix:         coeurSix,
		coeurSept:        coeurSept,
		coeurHuit:        coeurHuit,
		coeurNeuf:        coeurNeuf,
		coeurDix:         coeurDix,
		coeurValet:       coeurValet,
		coeurDame:        coeurDame,
		coeurRoi:         coeurRoi,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatalf("Erreur lors de l'exécution du jeu : %v", err)
	}
}

func main() {

	setImg()

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Casino de Léo")

}
