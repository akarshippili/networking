package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	port := 8080
	interface_ := fmt.Sprintf("127.0.0.1:%d", port)
	fmt.Printf("Starting a tcp Sever at %v!\n", interface_)

	ln, err := net.Listen("tcp", interface_)
	if err != nil {
		fmt.Printf("Error while starting the server, %v", err.Error())
		return
	}

	fmt.Println("Started server!!!")

	// start accepting connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			fmt.Printf("Error while accepting a connection, %v\n", err.Error())
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// conn.SetDeadline(time.Now().Add(10 * time.Second))

	bytes := make([]byte, 100)
	n, err := conn.Read(bytes)
	if err != nil {
		fmt.Printf("Error reading from connection %v, %v\n", conn, err.Error())
	}
	fmt.Printf("Received %v from conn\n", string(bytes[:n]))

	_, err = conn.Write(bytes[:n])
	if err != nil {
		fmt.Printf("Error writing to connection %v\n", err.Error())
		return
	}

	for {
		bytes = make([]byte, 100)
		_, err := conn.Write(time.Now().AppendFormat(bytes, time.RFC3339Nano))
		if err != nil {
			fmt.Printf("Error writing to connection %v\n", err.Error())
			return
		}
		time.Sleep(1 * time.Second)
	}
}
