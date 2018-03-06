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

// func boundaryCheck(chAbs chan mouselogger.Cords, lim int) {

// 	if chAbs.X >= lim-10 {
// 		fmt.Print("boundary")
// 	}

// }

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

func mouseRelTransmit(chRel chan mouselogger.Activity) {
	netcode.SendToActiveDevice(connectedDevices[activeDeviceIndex], port, chRel)
	wg.Done()
}

func main() {

	getRes()
	fmt.Println("Discovering clients on the network")
	connectedDevices = append(connectedDevices, netcode.GetOutboundIP())
	// connectedDevices = append(connectedDevices, netcode.DiscoverClients()...)
	fmt.Println("Done.")
	lim := getRes()
	chAbs := make(chan mouselogger.Cords)
	chAct := make(chan mouselogger.Activity)
	chKey := make(chan uint16)

	wg.Add(2)
	go mouseListener(chAbs, chAct)
	/*////////////////////////////////////////////////////////////
	I dont know how the channels are working
	so I made the function once see how the channel
	should be passed
	/////////////////////////////////*/
	go mouselogger.BoundaryCheck(chAbs, lim)
	// netcode.DiscoverClients()
	// go keebListener(ch_key)
	go mouseRelTransmit(chAct)

	wg.Wait()
	close(chAct)
	close(chAbs)
	close(chKey)
}

// func sendToActiveDevice(cords Cords) {
// 	cords.
// }
