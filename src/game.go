package src

import (
	"fmt"
	"reflect"
)

// Game ...
type Game struct {
	Board *Board

	Snake *Snake

	Food *Food

	Score int
	Round int
}

// NewGame ...
func NewGame(boardLength, boardWidth int) *Game {
	g := &Game{
		Board: newBoard(boardLength, boardWidth),
		Food:  &Food{},
		Score: 0,
		Round: 0,
	}
	g.spawnSnake()
	g.spawnFood()
	return g
}

// MoveSnake ...
func (g *Game) MoveSnake(newHeadPosition *Coordinates) error {
	if !g.Board.positionInBoardArea(newHeadPosition) {
		return fmt.Errorf("out of board")
	}
	if reflect.DeepEqual(newHeadPosition, g.Food.Position) {
		g.Snake.grow(g.Food)
		g.Score++
		g.Board.removeObject(g.Food.Position)
		g.spawnFood()
		return nil
	}
	g.moveSnakeParts(newHeadPosition)
	return nil
}

// PrintState ...
func (g *Game) PrintState() {
	fmt.Printf("\nBOARD AREA: %d X %d\nROUND: %d\nSCORE: %d\nSNAKE LENGTH: %d\nSNAKE HEAD COORDINATES: [%d, %d]\n",
		g.Board.length+1,
		g.Board.width+1,
		g.Round,
		g.Score,
		g.Snake.Lenght,
		g.Snake.Head.Position.X,
		g.Snake.Head.Position.Y,
	)
}

// Over ...
func (g *Game) Over(err error) {
	fmt.Printf("\n{{%s}}\n", err)
	fmt.Printf("\nGame over!\nFinal score: %d\n", g.Score)
}

func (g *Game) spawnSnake() {
	g.Snake = newSnake(g.Board.length, g.Board.width)
	_ = g.Board.RenderSnake(g.Snake)

}

func (g *Game) spawnFood() {
	for {
		f := newFood(g.Board.length, g.Board.width)
		if err := g.Board.setObject(f.Position, foodType); err != nil {
			continue
		}
		if reflect.DeepEqual(f.Position, g.Food.Position) {
			continue
		}
		g.Food = f
		break
	}
}

func (g *Game) moveSnakeParts(newHeadPosition *Coordinates) {
	g.Snake.TailProjection = g.Snake.Tail.Position
	part := g.Snake.Tail
	for {
		if part.NextPart != nil {
			part.Position = part.NextPart.Position
			part = part.NextPart
			continue
		}
		break
	}
	g.Snake.Head.Position = newHeadPosition
}
