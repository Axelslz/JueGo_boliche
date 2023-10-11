package scenes

import (
	"boliche/models"
	"boliche/views"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"fmt"
)

type GameScene struct {
	GameModel     *models.Game
	collisionChan chan bool
	loadedChan    chan bool
}

func NewGameScene() *GameScene {
	gs := &GameScene{
		GameModel: &models.Game{
			Ball: models.Ball{
				X:     380,
				Y:     550,
				SpeedY: 0,
			},
			Pins: []models.Pin{
				{X: 350, Y: 50, Position: models.Front},
				{X: 400, Y: 50, Position: models.BackLeft},
				{X: 300, Y: 50, Position: models.BackRight},
			},
			Running: true,
			Attempts: 3, // Añadido: Inicializar con 3 intentos.
		},
		collisionChan: make(chan bool),
		loadedChan:    make(chan bool),
	}

	go gs.CollisionThread()
	go views.LoadResources(gs.loadedChan)

	return gs
}

func (gs *GameScene) CollisionThread() {
	for range gs.collisionChan {
		for i := range gs.GameModel.Pins {
			if !gs.GameModel.Pins[i].Hit && bolaGolpeaPin(gs.GameModel.Ball, gs.GameModel.Pins[i]) {
				gs.GameModel.Pins[i].Hit = true
				gs.GameModel.Score++
			}
		}
	}
}

func (gs *GameScene) Update() error {
	select {
	case <-gs.loadedChan:
		fmt.Println("Espacio presionado!")
	default:
	}

	gs.MoveBall()

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		gs.GameModel.Ball.SpeedY = -5.0
	}

	gs.collisionChan <- true

	return nil
}

func (gs *GameScene) MoveBall() {
	gs.GameModel.Ball.SpeedX = 0

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		gs.GameModel.Ball.SpeedX = -3
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		gs.GameModel.Ball.SpeedX = 3
	}

	gs.GameModel.Ball.X += gs.GameModel.Ball.SpeedX
	gs.GameModel.Ball.Y += gs.GameModel.Ball.SpeedY

	if gs.GameModel.Ball.Y < 0 {
		gs.GameModel.Ball.Y = 550
		gs.GameModel.Ball.SpeedY = 0
		gs.GameModel.Attempts-- // Añadido: Reducir intentos

		// Añadido: Si no hay más intentos, detener el juego
		if gs.GameModel.Attempts <= 0 {
			gs.GameModel.Running = false
		}
	}
}

func (gs *GameScene) Draw(screen *ebiten.Image) {
	views.DrawGame(screen, gs.GameModel)
	views.DrawScore(screen, gs.GameModel.Score)
	views.DrawAttempts(screen, gs.GameModel.Attempts) 
	if !gs.GameModel.Running {
		views.DrawGameOver(screen)
	}
}

func (gs *GameScene) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}


func bolaGolpeaPin(bola models.Ball, pin models.Pin) bool {
    radioBola := 15.0  
    anchoPin := 20.0   
    altoPin := 40.0    

    // Revisa si el lado derecho de la bola es más a la izquierda que el lado izquierdo del pin
    if bola.X + radioBola < pin.X - anchoPin/2 {
        return false
    }
    // Revisa si el lado izquierdo de la bola es más a la derecha que el lado derecho del pin
    if bola.X - radioBola > pin.X + anchoPin/2 {
        return false
    }
    // Revisa si la parte inferior de la bola está más arriba que la parte superior del pin
    if bola.Y + radioBola < pin.Y - altoPin/2 {
        return false
    }
    // Revisa si la parte superior de la bola está más abajo que la parte inferior del pin
    if bola.Y - radioBola > pin.Y + altoPin/2 {
        return false
    }

    return true
}

