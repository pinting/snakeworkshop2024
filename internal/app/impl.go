package app

import (
	"github.com/gdamore/tcell/v2"
	"github.com/pinting/snakeworkshop2024/internal/common/drawing"
	"github.com/pinting/snakeworkshop2024/internal/common/random"
)

var snake []drawing.Coord
var food drawing.Coord

func begin() {
	snake = []drawing.Coord{{X: 2, Y: 2}}
	direction = drawing.Coord{X: 1, Y: 0}

	generateFood()
}

func tick() {
	step()
	drawSnake()
	drawFood()
	drawBorder()
}

func drawSnake() {
	for _, segment := range snake {
		put(segment, '■', tcell.ColorGreen)
	}
}

func drawFood() {
	put(food, '▣', tcell.ColorRed)
}

func isSnakeHit(p drawing.Coord) bool {
	for _, segment := range snake {
		if segment.X == p.X && segment.Y == p.Y {
			return true
		}
	}

	return false
}

func isSnakeKilled() bool {
	for i := 0; i < len(snake); i++ {
		for j := 0; j < len(snake); j++ {
			if i != j && snake[i].X == snake[j].X && snake[i].Y == snake[j].Y {
				return true
			}
		}
	}

	head := snake[len(snake)-1]

	return head.X < 1 || head.X >= size-1 || head.Y < 1 || head.Y >= size-1
}

func isFoodEaten() bool {
	head := snake[len(snake)-1]

	return head.X == food.X && head.Y == food.Y
}

func generateFood() {
	for {
		food = drawing.Coord{X: random.RandInt(1, size-2), Y: random.RandInt(1, size-2)}

		if !isSnakeHit(food) {
			break
		}
	}
}

func step() {
	if len(snake) == 0 {
		return
	}

	tail := snake[0]

	for i := 0; i < len(snake)-1; i++ {
		snake[i] = snake[i+1]
	}

	snake[len(snake)-1].X += direction.X
	snake[len(snake)-1].Y += direction.Y

	if isSnakeKilled() {
		begin()
	} else if isFoodEaten() {
		snake = append([]drawing.Coord{tail}, snake...)

		generateFood()
	}
}

func drawBorder() {
	for x := 0; x < size; x++ {
		put(drawing.Coord{X: x, Y: 0}, '-', tcell.ColorWhite)
		put(drawing.Coord{X: x, Y: size - 1}, '-', tcell.ColorWhite)
	}

	for y := 0; y < size; y++ {
		put(drawing.Coord{X: 0, Y: y}, '|', tcell.ColorWhite)
		put(drawing.Coord{X: size - 1, Y: y}, '|', tcell.ColorWhite)
	}

	put(drawing.Coord{X: 0, Y: 0}, '+', tcell.ColorWhite)
	put(drawing.Coord{X: size - 1, Y: 0}, '+', tcell.ColorWhite)
	put(drawing.Coord{X: 0, Y: size - 1}, '+', tcell.ColorWhite)
	put(drawing.Coord{X: size - 1, Y: size - 1}, '+', tcell.ColorWhite)
}
