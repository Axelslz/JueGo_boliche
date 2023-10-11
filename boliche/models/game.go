package models

type Game struct {
	Ball    Ball
	Pins    []Pin
	Running bool
	SpacePressed bool
	Score   int
	Attempts int
}

type PinPosition int

const (
	Front     PinPosition = iota
	BackLeft
	BackRight
)

type Pin struct {
	X, Y     float64
	Position PinPosition
	Hit      bool 
}