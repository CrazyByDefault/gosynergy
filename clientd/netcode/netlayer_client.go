package netcode

import (
	"bufio"
	"bytes"
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

	var network bytes.Buffer
	dec := gob.NewDecoder(&network)

	buf := make([]byte, 1024)

	for {
		var recievedItem mouselogger.Activity
		n, addr, err := ServerConn.ReadFromUDP(buf)
		dec.Decode(&recievedItem)
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

// ListenForHost waits for the initial host ping, and responds with a pong.
func ListenForHost() {
	ln, _ := net.Listen("tcp", ":6969")

	conn, _ := ln.Accept()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')

		if string(message) == "ping" {
			conn.Write([]byte("pong\n"))
			println("Connected to host")
			break
		}
	}
}
