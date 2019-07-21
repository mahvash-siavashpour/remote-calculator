package main

import (
	client "Client"
	server "Server"
)

func main() {

	server.Server_Run()
	client.Client_Run()
}
