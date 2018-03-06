package netcode

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"../mouselogger"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

// SendToActiveDevice sends stuff to the active device using UDP
func SendToActiveDevice(deviceIP net.IP, port int, chRel chan mouselogger.Activity) {

	list := []string{deviceIP.String(), ":", strconv.Itoa(port)}
	var ServerIP bytes.Buffer
	for _, l := range list {
		ServerIP.WriteString(l)
	}

	ServerAddr, err := net.ResolveUDPAddr("udp", ServerIP.String())
	checkError(err)

	Conn, err := net.DialUDP("udp", nil, ServerAddr)
	checkError(err)

	var network bytes.Buffer
	enc := gob.NewEncoder(&network)

	defer Conn.Close()

	for {
		item := <-chRel
		encErr := enc.Encode(item)
		if err != nil {
			log.Fatal("encode error:", encErr)
		}

		buf := network.Bytes()
		_, err := Conn.Write(buf)
		if err != nil {
			fmt.Println(item, err)
		}
		network.Reset()
		// time.Sleep(time.Second * 1)
	}
}

// GetOutboundIP gets the active IP of the own machine through magic :)
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

// DiscoverClients scans the network and returns all the active clients on the network
func DiscoverClients() []net.IP {
	var clients []net.IP

	arp := "arp | sed -n 's/.*\\(\\(\\(^\\| \\)[0-9]\\{1,3\\}\\.\\)\\{1\\}\\([0-9]\\{1,3\\}\\.\\)\\{2\\}[0-9]\\{1,3\\}\\) .*/\\1/gp'"
	cmd, _ := exec.Command("bash", "-c", arp).Output()
	// fmt.Print(arp)
	IPList := fmt.Sprintf("%s", cmd)

	IPs := strings.Split(IPList, "\n")

	for _, IP := range IPs {
		if isClient(net.ParseIP(IP)) {
			clients = append(clients, net.ParseIP(IP))
		}
	}
	return clients
}

func isClient(IPToCheck net.IP) bool {
	conn, _ := net.Dial("tcp", IPToCheck.String()+":6969")
	fmt.Fprintf(conn, "ping\n")

	message, _ := bufio.NewReader(conn).ReadString('\n')
	if message == "pong" {
		return true
	}
	return false
}
