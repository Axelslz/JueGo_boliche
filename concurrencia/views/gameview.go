package views

import (
	"concurrencia/models"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

var carImage *ebiten.Image
var fondoImage *ebiten.Image

func init() {
    var err error
    fondoImage, _, err = ebitenutil.NewImageFromFile("./assets/background.png")
    if err != nil {
        log.Fatalf("Error al cargar la imagen de fondo: %v", err)
    }
}

func InitAssets() {
	// Cargar imágenes (para simplificar, estamos omitiendo errores)
	carImage, _, _ = ebitenutil.NewImageFromFile("assets/car.png")
}

func DrawCar(screen *ebiten.Image, car *models.Car) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(car.X, car.Y)
	screen.DrawImage(carImage, opts)
}

