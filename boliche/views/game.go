package views

import (
    "boliche/models"
    "github.com/hajimehoshi/ebiten/v2"
    // ... otros imports
)

func DrawGame(screen *ebiten.Image, game *models.Game) {
	// Dibuja el fondo escalado
	fondoOpts := &ebiten.DrawImageOptions{}
	fw, fh := fondoImage.Size()
	fondoOpts.GeoM.Scale(float64(screenWidth)/float64(fw), float64(screenHeight)/float64(fh))
	screen.DrawImage(fondoImage, fondoOpts)

	// Dibuja todos los pines
	for _, pin := range game.Pins {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Scale(0.2, 0.2)
		
		// Ajuste de posición basado en la posición inicial del pin
		switch pin.Position {
		case models.Front:
			opts.GeoM.Translate(pin.X, pin.Y)
		case models.BackLeft:
			opts.GeoM.Translate(pin.X - 40*0.2, pin.Y + 64*0.2)
		case models.BackRight:
			opts.GeoM.Translate(pin.X + 40*0.1, pin.Y + 64*0.1)
		}
		
		screen.DrawImage(pinImage, opts)
	}

	// Ajusta la posición de la bola para que esté en la parte inferior
    ballY := screenHeight - float64(ballImage.Bounds().Dy())/10.0
	ballX := screenWidth / 2.0

    // Dibuja la bola en la parte inferior
    opts := &ebiten.DrawImageOptions{}
    opts.GeoM.Scale(0.1, 0.1)
    opts.GeoM.Translate(ballX, ballY)
    screen.DrawImage(ballImage, opts)
	
}


func DrawBall(screen *ebiten.Image, ball *models.Ball) {
    // Utiliza la información de la bola para dibujarla en la posición correcta en la pantalla
    opts := &ebiten.DrawImageOptions{}
    opts.GeoM.Scale(0.1, 0.1) // Ajusta el factor de escala según tus necesidades
    opts.GeoM.Translate(ball.X, ball.Y)
    screen.DrawImage(ballImage, opts)  // Aquí deberías tener una imagen de la bola cargada previamente
}