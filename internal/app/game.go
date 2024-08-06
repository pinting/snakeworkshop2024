package app

import (
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/pinting/snakeworkshop2024/internal/common/client"
	"github.com/pinting/snakeworkshop2024/internal/common/drawing"
)

var direction drawing.Coord

func put(p drawing.Coord, r rune, c tcell.Color) {
	drawing.Put(p.X, p.Y, r, int(c))
	client.Queue(p.X, p.Y, int(r), int(c))
}

func onEventKey(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyUp:
		if direction.Y != 1 {
			direction.X = 0
			direction.Y = -1
		}
	case tcell.KeyLeft:
		if direction.X != 1 {
			direction.X = -1
			direction.Y = 0
		}
	case tcell.KeyDown:
		if direction.Y != -1 {
			direction.X = 0
			direction.Y = 1
		}
	case tcell.KeyRight:
		if direction.X != -1 {
			direction.X = 1
			direction.Y = 0
		}
	case tcell.KeyEscape:
		drawing.Fini()
		os.Exit(0)
	}

	return event
}

func loop() {
	ticker := time.NewTicker(speed)

	defer ticker.Stop()

	for range ticker.C {
		drawing.Clear()
		tick()
		client.Broadcast()
		drawing.Show()
	}
}
