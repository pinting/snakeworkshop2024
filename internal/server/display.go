package server

import (
	"fmt"

	"github.com/pinting/snakeworkshop2024/internal/common/drawing"
	"github.com/pinting/snakeworkshop2024/internal/server/virtualdisplay"
)

var p = drawing.Coord{X: 0, Y: 0}

const displayW = 400
const displayH = 50

var step = drawing.Coord{X: 10, Y: 10}

func display() {
	drawing.Clear()

	for x := 0; x < displayW; x++ {
		for y := 0; y < displayH; y++ {
			x2 := p.X + x
			y2 := p.Y + y
			du := virtualdisplay.Get(x2, y2)

			drawing.Put(x, y+1, du.R, du.C)
		}
	}

	drawing.Print(fmt.Sprintf("%d %d", p.X, p.Y), drawing.Put)
	drawing.Show()
}

func moveUp() {
	p.Y -= step.Y
	p.Y = max(0, p.Y)
}

func moveDown() {
	h := virtualdisplay.VirtualDisplayH - displayH
	p.Y += step.Y
	p.Y = min(h, p.Y)
}

func moveRight() {
	w := virtualdisplay.VirtualDisplayW - displayW
	p.X += step.X
	p.X = min(w, p.X)
}

func moveLeft() {
	p.X -= step.X
	p.X = max(0, p.X)
}
