package netcode

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
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

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
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
		time.Sleep(time.Second * 1)
	}
}
