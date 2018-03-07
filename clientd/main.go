package main

import (
	"sync"

	"./netcode"
	"./mousemover"
)

var wg sync.WaitGroup

func mouseInputListener(inChan chan mousemover.Activity) {
	go func() {
		for {
			mousemover.ReadMouse(inChan)
		}
	}()
	wg.Done()
}

func main() {
	var mouseChan chan mousemover.Activity

	netcode.ListenForHost()

	wg.Add(2)
	go mouseInputListener(mouseChan)
	netcode.RecieveFromHost(mouseChan)

	wg.Wait()
}
