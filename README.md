# Simple console snake

&nbsp;

## Installation && compilation

      git clone git@github.com:bestpilotingalaxy/sky-snake.git
      
      cd sky-snake/cmd


---
      go build -o snake-game
      
### Run      
      ./snake-game
      
      
&nbsp;

## State output example

```
BOARD AREA: 10 X 30
ROUND: 173
SCORE: 12
SNAKE LENGTH: 14
SNAKE HEAD COORDINATES: [6, 23]
```

&nbsp;

## Commands

```
W - up
A - left
S - down
D - right
```
&nbsp;

## Layout

```
.
├── cmd
│   ├── cli.go
│   └── main.go
├── go.mod
├── README.md
└── src
    ├── board.go
    ├── coordinates.go
    ├── food.go
    ├── game.go
    └── snake.go
```
