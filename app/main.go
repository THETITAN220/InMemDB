package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("*** SERVER STARTING ***")

	l, err := net.Listen("tcp", ":6379")

	if err != nil {
		fmt.Println("Failed to bind to the port 6379")
		os.Exit(1)
	}

	fmt.Println("Server listening on port 6379. Waiting for connections....")

	for {
		conn, err := l.Accept()

		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

	defer conn.Close()

	fmt.Printf("Client connected to the port %s\n", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "PING" {
			_, err := conn.Write([]byte("+PONG\r\n"))
			if err != nil {
				fmt.Println("Error writing PONG response: ", err.Error())
				break
			}
		} else {
			errorMessage := fmt.Sprintf(" Error invalid command: %s\r\n ", text)
			_, err := conn.Write([]byte(errorMessage))
			if err != nil {
				fmt.Println("Error writing error response: ", err.Error())
				break
			}
		}

		if err := scanner.Err(); err != nil && err.Error() != "EOF" {
			fmt.Printf("Error reading from connection %s: %v \n", conn.RemoteAddr(), err)
		}

	}

	fmt.Printf("Client disconnected from %s \n", conn.RemoteAddr())
}
