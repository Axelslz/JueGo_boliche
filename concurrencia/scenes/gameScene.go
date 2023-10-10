package scenes

import (
	"concurrencia/models"
	"concurrencia/views"
	"github.com/hajimehoshi/ebiten/v2"
)

var fondoImage *ebiten.Image

type GameScene struct {
	Car *models.Car
}

func NewGameScene() *GameScene {
	return &GameScene{
		Car: &models.Car{
			X: 50,
			Y: 100,
			Velocity: 0,
		},
	}
}

func (gs *GameScene) Update() error {
	// Lógica de actualización, como mover el auto, aquí.
	return nil
}

func (gs *GameScene) Draw(screen *ebiten.Image) {
	screen.DrawImage(fondoImage, &ebiten.DrawImageOptions{})
	views.DrawCar(screen, gs.Car)
}

func (gs *GameScene) Layout(w, h int) (int, int) {
	return 320, 240
}


