package views

import (
	"boliche/models"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"sync"
)

const (
	screenWidth  = 800.0
	screenHeight = 600.0
)

var (
	pinImage   *ebiten.Image
	fondoImage *ebiten.Image
	ballImage  *ebiten.Image
)

// LoadResources carga los recursos de imágenes en goroutines y utiliza un canal para comunicar el estado.
func LoadResources(loadedChan chan bool) {
	var wg sync.WaitGroup
	wg.Add(3) 

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
