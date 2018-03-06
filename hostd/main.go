package main

import (
	"fmt"
	"net"
	"os/exec"

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

func main() {
	connectedDevices = append(connectedDevices, netcode.GetOutboundIP())
	netcode.SendToActiveDevice(connectedDevices[activeDeviceIndex], port)
}
