package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

type GameState int

const (
	Menu GameState = iota
	Roulette
	Blackjack
)

type Game struct {
	img           *ebiten.Image
	ballImg       *ebiten.Image
	angle         float64
	ballAngle     float64
	ballRadius    float64
	state         GameState
	rouletteStart time.Time
	isSpinning    bool
	ballVisible   bool
}

var rouletteNumbers = [37]int{
	0, 32, 15, 19, 4, 21, 2, 25, 17, 34, 6, 27, 13, 36, 11, 30, 8, 23, 10, 5, 24, 16, 33, 1, 20, 14, 31, 9, 22, 18, 29, 7, 28, 12, 35, 3, 26,
}

func (g *Game) Update() error {
	if g.state == Menu {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			if x >= 220 && x <= 320 && y >= 220 && y <= 260 {
				g.state = Blackjack
			} else if x > 320 && x <= 420 && y >= 220 && y <= 260 {
				g.state = Roulette
			}
		}
	} else if g.state == Roulette {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			if x >= 220 && x <= 320 && y >= 400 && y <= 440 {
				g.isSpinning = true
				g.ballVisible = true
				g.rouletteStart = time.Now()
				g.ballAngle = rand.Float64() * 2 * math.Pi // Random starting angle
				g.ballRadius = 90.0                        // Set a fixed radius for the ball's path in the outer circle
			}
		}
		if g.isSpinning {
			g.angle += 0.05
			g.ballAngle += 0.1

			if time.Since(g.rouletteStart).Seconds() > 10 {
				g.isSpinning = false
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
	}
}

func (g *Game) drawMenu(screen *ebiten.Image) {
	msg := "Choisissez votre jeu : "
	text.Draw(screen, msg, basicfont.Face7x13, 240, 200, color.Black)

	// Draw buttons
	ebitenutil.DrawRect(screen, 220, 220, 100, 40, color.RGBA{170, 170, 0, 255})
	ebitenutil.DrawRect(screen, 320, 220, 100, 40, color.RGBA{170, 0, 170, 255})

	// Draw button text
	text.Draw(screen, "Blackjack", basicfont.Face7x13, 235, 245, color.Black)
	text.Draw(screen, "Roulette", basicfont.Face7x13, 340, 245, color.Black)
}

func (g *Game) drawBlackjack(screen *ebiten.Image) {

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

	// Draw button
	ebitenutil.DrawRect(screen, 220, 400, 100, 40, color.RGBA{170, 170, 170, 255})
	text.Draw(screen, "Lancer une balle", basicfont.Face7x13, 225, 425, color.Black)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func main() {
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

	game := &Game{
		img:     img,
		ballImg: ballImg,
		state:   Menu,
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Casino de Léo")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatalf("Erreur lors de l'exécution du jeu : %v", err)
	}
}
