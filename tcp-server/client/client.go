package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error while connecting to server", err.Error())
		return
	}
	defer conn.Close()

	if len(os.Args) <= 1 {
		return
	}

	message := strings.Join(os.Args[1:], " ")
	_, err = conn.Write(([]byte)(message))
	if err != nil {
		fmt.Println("Error while writing to connection", err.Error())
	}

	for {
		bytes := make([]byte, 100)
		n, err := conn.Read(bytes)
		if err != nil {
			fmt.Printf("Error reading from connection %v, %v\n", conn, err.Error())
			return
		}
		fmt.Printf("Received: %v\n", string(bytes[:n]))
	}
}
