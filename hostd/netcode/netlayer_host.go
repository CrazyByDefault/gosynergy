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
	"time"

	"../mouselogger"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

// ConnectToActiveDevice establishes a UDP connection with active device
func ConnectToActiveDevice(deviceIP net.IP, port int) net.Conn {
	list := []string{deviceIP.String(), ":", strconv.Itoa(port)}
	var ServerIP bytes.Buffer
	for _, l := range list {
		ServerIP.WriteString(l)
	}

	ServerAddr, err := net.ResolveUDPAddr("udp", ServerIP.String())
	checkError(err)

	Conn, err := net.DialUDP("udp", nil, ServerAddr)
	checkError(err)
	fmt.Println("Connected")
	return Conn
}

// SendToActiveDevice sends mouse activity to the active device using UDP
func SendToActiveDevice(Conn net.Conn, chRel chan mouselogger.Activity, chSwitch chan bool) {
	fmt.Println("Sending?")
	enc := gob.NewEncoder(Conn)
	defer Conn.Close()

	for {
		select {
		default:
			fmt.Println("Sending.")
			item := <-chRel
			fmt.Print(item)
			encErr := enc.Encode(item)
			if encErr != nil {
				log.Fatal("encode error:", encErr)
			}
			break

		case <-chSwitch:
			fmt.Println("Closing connection")
			for len(chSwitch) > 0 {
				<-chSwitch
			}
			Conn.Close()
			return
		}
		// time.Sleep(time.Second * 1)
	}
}

// CloseActiveDevice closes the connection to the device
func CloseActiveDevice(Conn net.Conn) {
	Conn.Close()
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

	fmt.Println("Checking if " + IPToCheck.String() + " is a client")

	conn, err := net.DialTimeout("tcp", IPToCheck.String()+":8080", 1*time.Second)
	fmt.Println("Dialed up")

	if err != nil {
		// log.Print("Error: ", err)
		return false
	}
	defer conn.Close()

	conn.Write([]byte("ping\r\n\r\n"))

	conn.SetDeadline(time.Now().Add(3 * time.Second))
	buff := make([]byte, 1024)
	fmt.Println("Waiting for reply")
	n, _ := conn.Read(buff)
	if fmt.Sprintf("%s", buff[:n]) != "" {
		fmt.Println("Pong!")
		return true
	}

	return false
}

// ListenForHostBoundary waits for the initial host ping, and responds with a pong.
func ListenForClientBoundary(chSwitch chan bool) {
	ln, _ := net.Listen("tcp", ":8082")

	conn, _ := ln.Accept()
	fmt.Println("Client Boundary Listener Started")

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')

		if string(message) != "" {
			conn.Write([]byte("pong\r\n\r\n"))
			println("Back to host")
			chSwitch <- true
			return
		}
	}
}
