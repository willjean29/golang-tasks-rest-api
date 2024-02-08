package main

import (
	server "app/src/shared/infra"
)

func main() {
	newServer, _ := server.NewServer()
	newServer.Run()
}
