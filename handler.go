package main

import (
	"bufio"
	"io"
	"net"
)

//Handler is a loop dispatcher
type Handler struct {
	processor func(input string) string
}

// Handles the incoming connection
func (handler *Handler) handle(conn net.Conn) {

	defer conn.Close()

	// Read input from connection, keep it open until client exits
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {

		io.WriteString(conn, handler.processor(scanner.Text()))
	}
}
