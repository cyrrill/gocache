package main

import (
	"bufio"
	"io"
	"net"
)

// ScanHandler is a loop dispatcher
type ScanHandler struct {
	processor func(input string) string
}

// Handles the incoming connection
func (handler ScanHandler) handle(conn net.Conn) {

	// Close connection on handler exit
	defer conn.Close()

	// Setup scanner which reads connection input
	scanner := bufio.NewScanner(conn)

	// Loop while input on scanner
	for scanner.Scan() {

		// Read from network
		input := scanner.Text()

		// Dispatch processor function and get output
		output := handler.processor(input)

		// Transmit result to client
		io.WriteString(conn, output)
	}
}
