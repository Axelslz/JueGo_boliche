package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 4800   
	tileSize     = 32
)

var (
	playerX, playerY int
	foodX, foodY     int
	score            int
	gameOver         bool
	openMouth        bool
	keyState         map[ebiten.Key]bool
)

func init() {
	rand.Seed(time.Now().UnixNano())
	resetGame()
	keyState = make(map[ebiten.Key]bool)
}

func resetGame() {
	playerX, playerY = screenWidth/2, screenHeight/2
	foodX, foodY = rand.Intn(screenWidth/tileSize)*tileSize, rand.Intn(screenHeight/tileSize)*tileSize
	score = 0
	gameOver = false
	openMouth = false
}

type Game struct{}

func (g *Game) Update() error {
	if keyState[ebiten.KeyUp] {
		playerY -= tileSize
	}
	if keyState[ebiten.KeyDown] {
		playerY += tileSize
	}
	if keyState[ebiten.KeyLeft] {
		playerX -= tileSize
	}
	if keyState[ebiten.KeyRight] {
		playerX += tileSize
	}

	if playerX == foodX && playerY == foodY {
		score++
		foodX, foodY = rand.Intn(screenWidth/tileSize)*tileSize, rand.Intn(screenHeight/tileSize)*tileSize
	}

	if playerX < 0 || playerX >= screenWidth || playerY < 0 || playerY >= screenHeight {
		gameOver = true
	}

	if gameOver {
		if keyState[ebiten.KeySpace] {
			resetGame()
		}
	}

	if keyState[ebiten.KeyQ] {
		return fmt.Errorf("quit game")
	}

	// Cambiar el estado de la boca cada vez que Pac-Man se mueve
	openMouth = !openMouth

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0, 0, 0, 255})

	if !gameOver {
		// Dibujar Pac-Man
		pacManImg := ebiten.NewImage(tileSize, tileSize)
		pacManImg.Fill(color.NRGBA{255, 0, 0, 255})

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(playerX), float64(playerY))
		if openMouth {
			screen.DrawImage(pacManImg.SubImage(image.Rect(tileSize/4, tileSize/4, 3*tileSize/4, 3*tileSize/4)).(*ebiten.Image), op)
		} else {
			screen.DrawImage(pacManImg, op)
		}

		// Dibujar comida
		foodImg := ebiten.NewImage(tileSize, tileSize)
		foodImg.Fill(color.NRGBA{0, 255, 0, 255})
		screen.DrawImage(foodImg, &ebiten.DrawImageOptions{
			GeoM: ebiten.TranslateGeo(float64(foodX), float64(foodY)),
		})
	}

	if gameOver {
		msg := "Game Over! Score: Press Space to Restart"
		textWidth := len(msg) * 12
		ebitenutil.DebugPrintAt(screen, msg, (screenWidth-textWidth)/2, screenHeight/2)
	}

	scoreMsg := fmt.Sprintf("Score: %d", score)
	ebitenutil.DebugPrintAt(screen, scoreMsg, 10, 10)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Pac-Man Lite")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func updateKeyState(keys []ebiten.Key, state bool) {
	for _, key := range keys {
		keyState[key] = state
	}
}
