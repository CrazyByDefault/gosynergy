package main

import (
	"sync"

	"./mousemover"
	"./netcode"
)

var wg sync.WaitGroup

func mouseInputListener(inChan chan mousemover.Activity) {
	go func() {
		var ls, rs, ms int
		for {
			current := <-inChan
			ls, rs, ms = mousemover.ReadMouse(current, ls, rs, ms)
		}
	}()
	wg.Done()
}

func main() {
	mouseChan := make(chan mousemover.Activity)

	netcode.ListenForHost()

	wg.Add(2)
	go mouseInputListener(mouseChan)
	go netcode.RecieveFromHost(mouseChan)

	wg.Wait()
}
