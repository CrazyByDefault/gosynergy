package main

import (
	"./netcode"
)

func main() {
	netcode.ListenForHost()
	netcode.RecieveFromHost()
}
