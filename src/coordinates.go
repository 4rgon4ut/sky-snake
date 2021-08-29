package src

import "math/rand"

// Coordinates ...
type Coordinates struct {
	X, Y int
}

// NewPosition ...
func NewPosition(x, y int) *Coordinates {
	return &Coordinates{
		X: x,
		Y: y,
	}
}

// NewRandomPosition ...
func newRandomPosition(x, y int) *Coordinates {
	return &Coordinates{
		X: rand.Intn(x),
		Y: rand.Intn(y),
	}
}
