package netcode

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

// SendToActiveDevice sends stuff to the active device using UDP
func SendToActiveDevice(deviceIP net.IP, port int) {

	list := []string{deviceIP.String(), ":", strconv.Itoa(port)}
	var ServerIP bytes.Buffer
	for _, l := range list {
		ServerIP.WriteString(l)
	}

	ServerAddr, err := net.ResolveUDPAddr("udp", ServerIP.String())
	checkError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:7000")
	checkError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	checkError(err)

	defer Conn.Close()
	i := 0
	for {
		msg := strconv.Itoa(i)
		i++
		buf := []byte(msg)
		_, err := Conn.Write(buf)
		if err != nil {
			fmt.Println(msg, err)
		}
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
