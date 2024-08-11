package app

import (
	. "github.com/gdamore/tcell/v2"
	. "github.com/pinting/snakeworkshop2024/internal/common/drawing"
)

func begin() {

}

var blink = false

func tick() {
	gate := false

	if !gate {
		put(Coord{X: 0, Y: 1}, 's', ColorRed)
		put(Coord{X: 0, Y: 2}, 't', ColorRed)
		put(Coord{X: 0, Y: 3}, 'o', ColorRed)
		put(Coord{X: 0, Y: 4}, 'p', ColorRed)
	} else {
		for i, c := range "hello" {
			put(Coord{X: i, Y: 4}, c, ColorWhite)
		}

		for i := 0; i < 5; i++ {
			put(Coord{X: 6, Y: i}, rune('a'+i), ColorRed)
		}
	}

	if blink {
		put(Coord{X: 8, Y: 0}, '■', ColorRed)
		blink = false
	} else {
		put(Coord{X: 8, Y: 0}, '▣', ColorRed)
		blink = true
	}
}
