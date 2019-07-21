package main

import (
	_ "Client/main"
)

func main() {

	go Client.Client_Run()
}
