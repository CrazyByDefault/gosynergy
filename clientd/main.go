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
var activeDeviceIndex int

func getRes() int {
	toExec := "xdpyinfo  | grep -oP 'dimensions:\\s+\\K\\S+'"
	cmd, _ := exec.Command("bash", "-c", toExec).Output()
	s := strings.Split(string(cmd), "x")
	lim, _ := strconv.Atoi(s[0])

	return lim
}

func boundaryCheck(chAbs chan mouselogger.Cords, lim int, chSwitch chan bool) {
	go func() {
		for {
			current := <-chAbs
			if current.X >= lim-3 {
				fmt.Print("boundary")
				switchActiveDevice(chSwitch)
			}
		}
	}()
	wg.Done()
}

func switchActiveDevice(chSwitch chan bool) {
	if activeDeviceIndex == 0 {
		activeDeviceIndex = 1
	} else {
		activeDeviceIndex = 0
	}
	chSwitch <- true
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
	lim := getRes()
	chAbs := make(chan mouselogger.Cords)
	chSwitch := make(chan bool)

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
	go boundaryCheck(chAbs, lim, chSwitch)
	go func () {
		for {
			select {
			case <-chSwitch:
				netcode.ReturnToHost(host)
			}
		}
	}

	wg.Wait()
	close(chAbs)
	close(chSwitch)
}
