package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pinting/snakeworkshop2024/internal/common/random"
)

type Unit struct {
	X int
	Y int
	R int
	C int
}

type Message struct {
	Units []Unit
	User  string
}

const ServerAddress = "localhost:9999"

var User = fmt.Sprintf("%d", random.RandInt(1, 1000*1000))

var transmission []Unit

func Broadcast() {
	message := Message{
		Units: transmission,
		User:  User,
	}

	buffer := &bytes.Buffer{}
	err := json.NewEncoder(buffer).Encode(message)

	if err != nil {
		return
	}

	transmission = make([]Unit, 0)

	go func(buffer *bytes.Buffer) {
		url := fmt.Sprintf("http://%s", ServerAddress)
		resp, err := http.Post(url, "application/json", buffer)

		if err != nil {
			return
		}

		resp.Body.Close()
	}(buffer)
}

func Queue(x int, y int, r int, c int) {
	transmission = append(transmission, Unit{X: x, Y: y, R: r, C: c})
}
