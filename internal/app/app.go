package app

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/pinting/snakeworkshop2024/internal/common/drawing"
)

const size = 30
const speed = 150 * time.Millisecond

func Run() {
	begin()
	drawing.Loop(func() {
		go loop()
	}, func(ek *tcell.EventKey) {
		onEventKey(ek)
	})
}
