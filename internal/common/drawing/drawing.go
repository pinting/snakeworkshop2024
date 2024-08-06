package drawing

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

const debug = false

var screen tcell.Screen

type Coord struct {
	X int
	Y int
}

func Prepare() {
	var err error

	screen, err = tcell.NewScreen()

	if err != nil {
		print(err.Error())
		os.Exit(1)
	}

	if err = screen.Init(); err != nil {
		print(err.Error())
		os.Exit(1)
	}
}

func Put(x int, y int, r rune, c int) {
	screen.SetContent(x, y, r, nil, tcell.StyleDefault.Foreground(tcell.Color(c)))
}

func Fini() {
	screen.Fini()
}

func Clear() {
	screen.Clear()
}

func Show() {
	screen.Show()
}

func Loop(startLoop func(), onEventKey func(*tcell.EventKey)) {
	Prepare()

	defer Fini()

	Clear()
	Show()

	startLoop()

	for {
		event := screen.PollEvent()

		switch ev := event.(type) {
		case *tcell.EventKey:
			onEventKey(ev)
		case *tcell.EventResize:
			screen.Sync()
		}
	}
}

func Debug(format string, a ...any) (n int, err error) {
	if !debug {
		return 0, nil
	}

	return fmt.Printf(format, a...)
}

func Print(message string, put func(int, int, rune, int)) {
	const color = tcell.ColorWhite

	Debug("%s", message)

	const lineLength = 50

	for i, c := range message {
		x := i % lineLength
		y := (i - x) / lineLength

		put(x, y, c, int(color))
	}

}
