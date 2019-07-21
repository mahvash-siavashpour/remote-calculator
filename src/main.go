package main

import (
	client "Client"
	server "Server"
)

func main() {

	go server.RunServer()
	go client.RunClient()
}
