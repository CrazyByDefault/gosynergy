package main

import (
	"sync"

	"./mousemover"
	"./netcode"
)

var wg sync.WaitGroup

func mouseInputListener(inChan chan mousemover.Activity) {
	go func() {
		for {
			current := <-inChan
			mousemover.ReadMouse(current)
		}
	}()
	wg.Done()
}

func main() {
	var mouseChan chan mousemover.Activity

	netcode.ListenForHost()

	wg.Add(1)
	go mouseInputListener(mouseChan)
	go netcode.RecieveFromHost(mouseChan)

	wg.Wait()
}
