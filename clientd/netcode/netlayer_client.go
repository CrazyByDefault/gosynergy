package netcode

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"

	"../mouselogger"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

// RecieveFromHost recieves the data from the host -_-
func RecieveFromHost() {
	/* Lets prepare a address at any address at port 10001*/
	ServerAddr, err := net.ResolveUDPAddr("udp", ":7000")
	checkError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	checkError(err)
	defer ServerConn.Close()

	dec := gob.NewDecoder(ServerConn)

	for {
		var recievedItem mouselogger.Activity
		dec.Decode(&recievedItem)
		fmt.Println("Received")
		fmt.Print(recievedItem)

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
