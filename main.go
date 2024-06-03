package main

import (
	server "app/src/shared/infra"
)

func main() {
	// server with gorilla/mux
	// newServer, _ := server.NewServer()
	// newServer.Run()

	// server with gin
	newServer := server.CreateServer()
	newServer.Run()
}
