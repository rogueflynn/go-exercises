package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	var d net.Dialer
	buffer := make([]byte, 1024)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	reader := bufio.NewReader(os.Stdin)
	defer cancel()

	for true {
		conn, err := d.DialContext(ctx, "tcp", "localhost:8080")
		fmt.Printf("Your Message: ")
		message, _ := reader.ReadString('\n')
		if _, err := conn.Write([]byte(message)); err != nil {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatalf("failed to dial: %v", err)
		}
		conn.Read(buffer)
		response := string(buffer)
		// Remove any NULL characters from 'b'
		fmt.Println(response)
	}
}
