package server

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/pinting/snakeworkshop2024/internal/common/client"
	"github.com/pinting/snakeworkshop2024/internal/common/drawing"
	"github.com/pinting/snakeworkshop2024/internal/server/virtualdisplay"
)

func root(w http.ResponseWriter, req *http.Request) {
	if req.Body == nil {
		drawing.Debug("empty body")
		return
	}

	var message client.Message

	err := json.NewDecoder(req.Body).Decode(&message)

	if err != nil {
		drawing.Debug(err.Error())
		return
	}

	process(&message)
}

func process(message *client.Message) {
	entry := virtualdisplay.MakeEntry(message.User)

	entry.LastActivity = time.Now().Unix()
	entry.Units = message.Units
}

func loop() {
	for {
		virtualdisplay.Render()
		display()
	}
}

func listenAndServe() {
	http.HandleFunc("/", root)
	http.ListenAndServe(client.ServerAddress, nil)
}

func Run() {
	virtualdisplay.Setup()
	drawing.Loop(func() {
		go listenAndServe()
		go loop()
	}, func(ek *tcell.EventKey) {
		switch ek.Key() {
		case tcell.KeyLeft:
			moveLeft()
		case tcell.KeyRight:
			moveRight()
		case tcell.KeyUp:
			moveUp()
		case tcell.KeyDown:
			moveDown()
		case tcell.KeyEscape:
			drawing.Fini()
			os.Exit(0)
		}
	})
}
