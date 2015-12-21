package main

import (
	"io"
	"net"
	"strings"
)

// PostHandler reads an HTTP POST
type PostHandler struct {
	processor func(input string) string
}

// Handles the incoming connection
func (handler PostHandler) handle(conn net.Conn) {

	// Close connection on handler exit
	defer conn.Close()

	b := make([]byte, 4096)

	// Read from network
	conn.Read(b)

	input := string(b)

	// Split Head from Body
	use := false
	var body string
	for _, val := range strings.Split(input, "\r\n") {
		if use {
			body = val
			break
		}
		if val == "" {
			use = true
		}
	}

	// Dispatch processor function and get output
	output := handler.processor(body)

	// Transmit result to client
	io.WriteString(conn, output)
}
