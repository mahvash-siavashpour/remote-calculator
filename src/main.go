package main

import (
	"client1"
	"server1"
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
