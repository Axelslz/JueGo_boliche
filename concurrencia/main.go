package main

import (
	"log"
	"concurrencia/scenes"
	"concurrencia/views"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// Inicializar assets como imágenes
	views.InitAssets()

	gameScene := scenes.NewGameScene()

	ebiten.SetWindowSize(320, 240)
	ebiten.SetWindowTitle("Race Game")

	if err := ebiten.RunGame(gameScene); err != nil {
		log.Fatal(err)
	}
}
