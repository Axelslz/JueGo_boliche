package main

import (
	"boliche/scenes"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	
	gameScene := scenes.NewGameScene()

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Boliche Game")

	// Ejecutar el juego
	if err := ebiten.RunGame(gameScene); err != nil {
		log.Fatal(err)
	}
}
