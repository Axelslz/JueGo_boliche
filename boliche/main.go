package main

import (
    "log"
    "github.com/hajimehoshi/ebiten/v2"
    "boliche/scenes"
)

func main() {
    gameScene := &scenes.GameScene{}

    ebiten.SetRunnableOnUnfocused(true)
    if err := ebiten.RunGame(gameScene); err != nil {
        log.Fatal(err)
    }
}



