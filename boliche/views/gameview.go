package views

import (
	"boliche/models"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
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

var ball models.Ball

func init() {
	var err error

	pinImage, _, err = ebitenutil.NewImageFromFile("assets/pin.png")
	if err != nil {
		log.Fatalf("Error al cargar imagen del pin: %v", err)
	}

	fondoImage, _, err = ebitenutil.NewImageFromFile("assets/background.png")
	if err != nil {
		log.Fatalf("Error al cargar imagen del fondo: %v", err)
	}

	ballImage, _, err = ebitenutil.NewImageFromFile("assets/ball.png")
	if err != nil {
		log.Fatalf("Error al cargar imagen de la bola: %v", err)
	}

    // Inicializa la posición y velocidad de la bola
    ball = models.Ball{
        X:      screenWidth / 2, // Posición inicial en el centro del ancho de la pantalla
        Y:      screenHeight - 20, // Posición inicial cerca de la parte inferior de la pantalla
        SpeedY: -5, // Velocidad hacia arriba (negativo para moverse hacia arriba)
    }
}

