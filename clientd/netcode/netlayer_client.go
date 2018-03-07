package netcode

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"../mousemover"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

// RecieveFromHost recieves and decodes the data from the host
func RecieveFromHost(inChan chan mousemover.Activity) {

	ServerAddr, err := net.ResolveUDPAddr("udp", ":7000")
	checkError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	checkError(err)
	defer ServerConn.Close()

	dec := gob.NewDecoder(ServerConn)

	for {
		var recievedItem mousemover.Activity
		dec.Decode(&recievedItem)
		fmt.Println("Received")
		fmt.Print(recievedItem)
		inChan <- recievedItem

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

// ListenForHost waits for the initial host ping, and responds with a pong.
func ListenForHost() net.IP {
	ln, _ := net.Listen("tcp", ":8080")

	conn, _ := ln.Accept()
	fmt.Println("Connection accepted")

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')

		if string(message) != "" {
			conn.Write([]byte("pong\r\n\r\n"))
			println("Connected to host")
			break
		}
	}
	return net.ParseIP(strings.Split(conn.RemoteAddr().String(), ":")[0])
}

func ReturnToHost(IPToCheck net.IP) {

	fmt.Println("Checking if " + IPToCheck.String() + " is a client")

	conn, err := net.DialTimeout("tcp", IPToCheck.String()+":8082", 1*time.Second)
	fmt.Println("Dialed up")

	if err != nil {
		// log.Print("Error: ", err)
		return
	}
	defer conn.Close()

	conn.Write([]byte("ping\r\n\r\n"))

	conn.SetDeadline(time.Now().Add(3 * time.Second))
	buff := make([]byte, 1024)
	fmt.Println("Waiting for reply")
	n, _ := conn.Read(buff)
	if fmt.Sprintf("%s", buff[:n]) != "" {
		fmt.Println("Pong!")

	}

}
