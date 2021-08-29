package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bestpilotingalaxy/snake-test/src"
)

func createGame() *src.Game {
	var length int
	var width int

	fmt.Println("\nSet board length :")
	_, err := fmt.Scanln(&length)
	if err != nil {
		log.Fatalln("Invalid input for board length")
	}
	fmt.Println("\nSet board width :")
	_, err = fmt.Scanln(&width)
	if err != nil {
		log.Fatalln("Invalid input for board width")
	}

	g := src.NewGame(length, width)
	g.Board.SquareRepresentation()
	return g
}

func run(g *src.Game) error {
	var input string
	for {
		//
		g.Round++

		fmt.Println("\nCOMMANDS: \nW - up\nA - left\nS - down\nD - right")

		fmt.Scanln(&input)
		input = strings.ToLower(input)
		var newHeadPosition *src.Coordinates
		switch input {

		case "w":
			newHeadPosition = src.NewPosition(g.Snake.Head.Position.X-1, g.Snake.Head.Position.Y)

		case "a":
			newHeadPosition = src.NewPosition(g.Snake.Head.Position.X, g.Snake.Head.Position.Y-1)

		case "s":
			newHeadPosition = src.NewPosition(g.Snake.Head.Position.X+1, g.Snake.Head.Position.Y)

		case "d":
			newHeadPosition = src.NewPosition(g.Snake.Head.Position.X, g.Snake.Head.Position.Y+1)

		default:
			fmt.Println("Invalid Input")
			return fmt.Errorf("Invalid input")
		}
		if err := g.MoveSnake(newHeadPosition); err != nil {
			return err
		}
		if err := g.Board.RenderSnake(g.Snake); err != nil {
			return err
		}
		g.Board.SquareRepresentation()
		g.PrintState()
	}
}
