package main

import (
	//"github.com/mahvash-siavashpour/remote-calculator/client"
	//"github.com/mahvash-siavashpour/remote-calculator/server"
	"client"
	"server"
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
