package main

import (
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	"./keylogger"
	"./mouselogger"
	"./netcode"
)

var connectedDevices []net.IP
var activeDeviceIndex int

var wg sync.WaitGroup

const port = 7000

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

func keebListener(chKey chan uint16) {
	devs, err := keylogger.NewDevices()
	if err != nil {
		fmt.Println(err)
		return
	}

	//keyboard device file, on your system it will be diffrent!
	rd := keylogger.NewKeyLogger(devs[4])

	for true {
		go keylogger.KeyActivity(chKey, rd)
		fmt.Print(int(<-chKey))
	}
}

func mouseListener(chAbs chan mouselogger.Cords, chAct chan mouselogger.Activity) {
	go func() {
		for {
			mouselogger.GetMouseAbs(chAbs)
		}
	}()

	go func() {
		for {
			mouselogger.GetMouseRel(chAct)
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

func mouseRelTransmit(chRel chan mouselogger.Activity, chSwitch chan bool) {
	for {
		switch activeDeviceIndex {
		case 1:
			conn := netcode.ConnectToActiveDevice(connectedDevices[activeDeviceIndex], port)
			netcode.SendToActiveDevice(conn, chRel, chSwitch)
		}
	}
	wg.Done()
}

func main() {

	fmt.Println("Discovering clients on the network")
	connectedDevices = append(connectedDevices, netcode.GetOutboundIP())
	connectedDevices = append(connectedDevices, netcode.DiscoverClients()...)
	fmt.Print("Done: ")
	fmt.Print(connectedDevices)
	lim := getRes()
	chAbs := make(chan mouselogger.Cords)
	chAct := make(chan mouselogger.Activity)
	chKey := make(chan uint16)
	chSwitch := make(chan bool)

	wg.Add(3)
	go mouseListener(chAbs, chAct)
	go boundaryCheck(chAbs, lim, chSwitch)
	// netcode.DiscoverClients()
	// go keebListener(ch_key)
	go mouseRelTransmit(chAct, chSwitch)

	wg.Wait()
	close(chAct)
	close(chAbs)
	close(chKey)
	close(chSwitch)
}

// func sendToActiveDevice(cords Cords) {
// 	cords.
// }
