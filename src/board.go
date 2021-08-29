package src

import (
	"fmt"
	"strings"
)

//
const (
	foodType  = 7
	snakeType = 1
)

// Board ...
type Board struct {
	area   [][]int
	length int
	width  int
}

// NewBoard ...
func newBoard(length, width int) *Board {
	area := make([][]int, length)
	for i := range area {
		area[i] = make([]int, width)
	}
	return &Board{
		area:   area,
		length: length - 1,
		width:  width - 1,
	}
}

// SquareRepresentation ...
func (b *Board) SquareRepresentation() {
	var toPrint string
	for _, row := range b.area {
		r := row
		toPrint = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(r)), ""), "[]")
		fmt.Println(toPrint)
	}
}

// RenderSnake ...
func (b *Board) RenderSnake(s *Snake) error {
	if b.area[s.Head.Position.X][s.Head.Position.Y] == snakeType {
		return fmt.Errorf("snake crossed body")
	}
	snakePart := s.Tail
	for i := 0; i < int(s.Lenght); i++ {
		_ = b.setObject(snakePart.Position, snakeType)
		snakePart = snakePart.NextPart
	}
	if s.TailProjection != nil {
		b.removeObject(s.TailProjection)
	}
	return nil
}

// SetObject ...
func (b *Board) setObject(position *Coordinates, objType int) error {
	if b.area[position.X][position.Y] == 0 {
		b.area[position.X][position.Y] = objType
		return nil
	}
	return fmt.Errorf("position [%d, %d] already taken by other object", position.X, position.Y)
}

// RemoveObject ...
func (b *Board) removeObject(position *Coordinates) {
	b.area[position.X][position.Y] = 0
}

// PositionInBoardArea ...
func (b *Board) positionInBoardArea(pos *Coordinates) bool {
	if pos.X > b.length || pos.X < 0 {
		return false
	}
	if pos.Y > b.width || pos.Y < 0 {
		return false
	}
	return true
}
