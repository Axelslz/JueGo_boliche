package scenes

import (
    "boliche/models"
    "boliche/views"
    "github.com/hajimehoshi/ebiten/v2"
	"fmt"
)

type GameScene struct {
    game models.Game
}

const (
    screenWidth  = 800 // puedes ajustar esto a tu valor preferido
    screenHeight = 600 // puedes ajustar esto a tu valor preferido
)


func (g *GameScene) Update() error {
    // Lógica de actualización

    if ebiten.IsKeyPressed(ebiten.KeySpace) && !g.game.Ball.Moving {
		fmt.Println("Espacio presionado")  // Añade esta línea
		g.game.Ball.SpeedY = -5
		g.game.Ball.Moving = true
	}
	

    // Aplicar gravedad a la bola
    gravity := 0.10
    g.game.Ball.SpeedY += gravity

    // Actualizar posición de la bola
	g.game.Ball.Y += g.game.Ball.SpeedY

    
    if g.game.Ball.Y > screenHeight {
        g.game.Ball.Y = screenHeight - 20 
        g.game.Ball.SpeedY = 0
        g.game.Ball.Moving = false
    }

	return nil
}

func (g *GameScene) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return 800, 600 
}

func (g *GameScene) Draw(screen *ebiten.Image) {
    views.DrawGame(screen, &g.game)
}