package src

// Snake ...
type Snake struct {
	Head           *bodyPart
	Tail           *bodyPart
	TailProjection *Coordinates
	Lenght         int
}

// BodyPart ...
type bodyPart struct {
	NextPart *bodyPart
	Position *Coordinates
}

// NewSnake .,.
func newSnake(boardLength, boardWidth int) *Snake {
	head := &bodyPart{
		NextPart: nil,
		Position: newRandomPosition(boardLength-1, boardWidth-1),
	}
	tail := &bodyPart{
		NextPart: head,
		Position: NewPosition(head.Position.X+1, head.Position.Y),
	}

	return &Snake{
		Head:           head,
		Tail:           tail,
		Lenght:         2,
		TailProjection: nil,
	}
}

// Grow ...
func (s *Snake) grow(f *Food) {
	newHead := &bodyPart{
		Position: f.Position,
		NextPart: nil,
	}
	// Set new snake head as an next element for old head
	s.Head.NextPart = newHead
	// Set new head
	s.Head = newHead
	s.Lenght++
}
