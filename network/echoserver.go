package main

import (
	"fmt"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)

	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err.Error())
	}

	conn.Write([]byte("Echo Server: " + string(buffer)))
	conn.Close()
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	fmt.Println("Listening on port 8080")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		go handleConnection(conn)
	}
}
