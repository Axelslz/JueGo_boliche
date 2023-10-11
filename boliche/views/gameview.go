package views

import (
	"boliche/models"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"sync"
	"fmt"
)

const (
	screenWidth  = 800.0
	screenHeight = 600.0
)

var (
	pinImage   *ebiten.Image
	fondoImage *ebiten.Image
	ballImage  *ebiten.Image
	gameOverImage *ebiten.Image
)


func LoadResources(loadedChan chan bool) {
	var wg sync.WaitGroup
	wg.Add(4) 

	go func() {
		defer wg.Done()
		var err error
		pinImage, _, err = ebitenutil.NewImageFromFile("assets/pin.png")
		if err != nil {
			log.Fatalf("Error al cargar imagen del pin: %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		var err error
		fondoImage, _, err = ebitenutil.NewImageFromFile("assets/background.png")
		if err != nil {
			log.Fatalf("Error al cargar imagen del fondo: %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		var err error
		ballImage, _, err = ebitenutil.NewImageFromFile("assets/ball.png")
		if err != nil {
			log.Fatalf("Error al cargar imagen de la bola: %v", err)
		}
	}()
	go func() {
		var err error
		gameOverImage, _, err = ebitenutil.NewImageFromFile("assets/game_over.png")
		if err != nil {
			log.Fatalf("Error al cargar imagen de Game Over: %v", err)
		}
		loadedChan <- true
	}()

	wg.Wait() // Espera a que las tres goroutines finalicen.
	loadedChan <- true 
}

func DrawGame(screen *ebiten.Image, game *models.Game) {
	DrawBackground(screen)
	DrawPins(screen, game)
	DrawBall(screen, &game.Ball)
}

func DrawBackground(screen *ebiten.Image) {
	fondoOpts := &ebiten.DrawImageOptions{}
	fw, fh := fondoImage.Size()
	fondoOpts.GeoM.Scale(float64(screenWidth)/float64(fw), float64(screenHeight)/float64(fh))
	screen.DrawImage(fondoImage, fondoOpts)
}

func DrawPins(screen *ebiten.Image, game *models.Game) {
	for _, pin := range game.Pins {
		if pin.Hit {
			continue
		}
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Scale(0.2, 0.2)
		opts.GeoM.Translate(pin.X, pin.Y)
		screen.DrawImage(pinImage, opts)
	}
}

func DrawBall(screen *ebiten.Image, ball *models.Ball) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(0.1, 0.1)
	opts.GeoM.Translate(ball.X, ball.Y)
	screen.DrawImage(ballImage, opts)
}

func DrawScore(screen *ebiten.Image, score int) {
    ebitenutil.DebugPrint(screen, fmt.Sprintf("Puntuación: %d", score))
}

func DrawAttempts(screen *ebiten.Image, attempts int) {
    ebitenutil.DebugPrint(screen, fmt.Sprintf("\n\nIntentos restantes: %d", attempts))
}

func DrawGameOver(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	w, h := gameOverImage.Size()

	x := (800 - float64(w)) / 2
	y := (600 - float64(h)) / 2

	opts.GeoM.Translate(x, y)
	screen.DrawImage(gameOverImage, opts)
}
