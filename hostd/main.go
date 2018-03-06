package main

import (
	"fmt"
	"net"
	"os/exec"
	"sync"

	"./keylogger"
	"./mouselogger"
	"./netcode"
)

var connectedDevices []net.IP
var activeDeviceIndex int

const port = 7000

func getRes() {
	toExec := "xdpyinfo  | grep -oP 'dimensions:\\s+\\K\\S+'"
	cmd, _ := exec.Command("bash", "-c", toExec).Output()
	fmt.Printf("%s", cmd)
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
	for true {
		go mouselogger.GetMouseAbs(chAbs)
		go mouselogger.GetMouseRel(chAct)
	}
}

func mouseRelTransmit(chRel chan mouselogger.Activity) {
	connectedDevices = append(connectedDevices, netcode.GetOutboundIP(), chRel)
	netcode.SendToActiveDevice(connectedDevices[activeDeviceIndex], port)
}

func main() {

	getRes()

	var wg sync.WaitGroup

	chAbs := make(chan mouselogger.Cords)
	chAct := make(chan mouselogger.Activity)
	chKey := make(chan uint16)

	wg.Add(1)
	go mouseListener(chAbs, chAct)
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
