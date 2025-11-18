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

	fmt.Printf("Client conncted from %s\n", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.TrimSpace(text) == "PING" {
			conn.Write([]byte("+PONG\r\n"))
		}
		conn.Write([]byte(" Error invalid command " + text + "\r\n"))
	}
}
