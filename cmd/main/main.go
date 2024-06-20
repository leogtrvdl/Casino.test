package main

import (
	"Casinotest/games"
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
	img              *ebiten.Image
	ballImg          *ebiten.Image
	rouletteTableImg *ebiten.Image
	background       *ebiten.Image
	dosCartes        *ebiten.Image
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
	trefleReine      *ebiten.Image
	trefleRoi        *ebiten.Image
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
	playerHand       []games.Card
	dealerHand       []games.Card
	blackjackStart   bool
	isPlayerSelected bool
	playerNumber     int
	players          []games.Player
	currentPlayer    int
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
			}
			if g.isPlayerSelected {
				//création + mélange des cartes
				deck := games.CreateDeck()
				games.Shuffle(deck)
				//creation des joueurs
				joueurs := games.CreatePlayers(g.playerNumber)
				croupier := games.CreateCroupier()
				//on donne une carte a chaque joueur
				for j := range g.playerNumber {
					games.GiveRandomCard(&joueurs[j], &deck)
				}
				for i := 0; i < 2; i++ {
					games.GiveRandomCard(&croupier, &deck)
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
	ballOp.GeoM.Translate(225, 270)
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

	/*dosCartes, _, err := ebitenutil.NewImageFromFile("assets/cartes/dos_cartes.png")
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

	trefleReine, _, err := ebitenutil.NewImageFromFile("assets/cartes/trefle/dame.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}

	trefleRoi, _, err := ebitenutil.NewImageFromFile("assets/cartes/trefle/roi.png")
	if err != nil {
		log.Fatalf("Erreur lors du chargement de l'image : %v", err)
		os.Exit(1)
	}*/

	game := &Game{
		img:              img,
		ballImg:          ballImg,
		rouletteTableImg: rouletteTableImg,
		state:            Menu,
		background:       background,
		/*dosCartes:        dosCartes,
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
		trefleReine:      trefleReine,
		trefleRoi:        trefleRoi,*/
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
