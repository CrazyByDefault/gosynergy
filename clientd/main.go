package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	"./mouselogger"
	"./mousemover"
	"./netcode"
)

var wg sync.WaitGroup
var activeDeviceIndex = 1

func boundaryCheck(chAbs chan mouselogger.Cords, lim int) {
	go func() {
		for {
			current := <-chAbs
			if current.X <= 10 {
				fmt.Print("boundary")
				switchActiveDevice()
			}
		}
	}()
	wg.Done()
}

func switchActiveDevice() {
	if activeDeviceIndex == 0 {
		activeDeviceIndex = 1
	} else {
		activeDeviceIndex = 0
	}
}

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
	chAbs := make(chan mouselogger.Cords)

	mouseChan := make(chan mousemover.Activity)

	host := netcode.ListenForHost()

	wg.Add(4)
	go mouseInputListener(mouseChan)
	go netcode.RecieveFromHost(mouseChan)
	go func() {
		for {
			mouselogger.GetMouseAbs(chAbs)
		}
	}()
	go boundaryCheck(chAbs, lim)
	go func() {
		for {
			if activeDeviceIndex == 0 {
				netcode.ReturnToHost(host)
			} else {
				time.Sleep(1 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
	close(chAbs)
}
