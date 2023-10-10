package models

type Game struct {
	Ball    Ball
	Pins    []Pin
	Running bool
	SpacePressed bool
}

