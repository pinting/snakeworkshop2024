package main

import (
	"flag"

	"github.com/pinting/snakeworkshop2024/internal/app"
	"github.com/pinting/snakeworkshop2024/internal/common/client"
	"github.com/pinting/snakeworkshop2024/internal/server"
)

func main() {
	serverFlag := flag.Bool("server", false, "Run the application as a server")
	userFlag := flag.String("user", "", "Username of the client")

	flag.Parse()

	if *userFlag != "" {
		client.User = *userFlag
	}

	if *serverFlag {
		server.Run()
	} else {
		app.Run()
	}
}
