package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// Listen for incoming connections
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	// Accept incoming connections
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// Create a new goroutine for each connection
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// Use a scanner to read messages from the client
	scanner := bufio.NewScanner(conn)

	// Continuously read messages from the client and print them
	for scanner.Scan() {
		msg := scanner.Text()
		fmt.Println(msg)

		// If the message is "quit", break the loop
		if strings.ToLower(msg) == "quit" {
			break
		}
	}

	// Close the connection when the loop is done
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	conn.Close()
}
