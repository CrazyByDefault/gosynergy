package netcode

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

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
