package models

type Ball struct {
    X, Y       float64
    SpeedX, SpeedY float64
    Moving     bool
}

func NewBall() *Ball {
    return &Ball{
        X: 0.0,
        Y: 0.0,
        // Inicializa otros campos si es necesario...
    }
}