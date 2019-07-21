package main

import (
	client "Client"
	server "Server"
	"sync"
)

var WG sync.WaitGroup

func main() {
	WG.Add(1)
	go server.RunServer()
	WG.Add(1)
	go client.RunClient()
	WG.Wait()
}
