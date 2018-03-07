package netcode

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"

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
func ListenForHost() {
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
}
