package models

type PinPosition int

const (
	Front     PinPosition = iota
	BackLeft
	BackRight
)

type Pin struct {
	X, Y     float64
	Position PinPosition
}