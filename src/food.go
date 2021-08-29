package src

// Food ...
type Food struct {
	Points   uint
	Position *Coordinates
}

// NewFood ...
func newFood(x, y int) *Food {
	return &Food{
		// Hardcoded points
		Points:   1,
		Position: newRandomPosition(x, y),
	}

}
