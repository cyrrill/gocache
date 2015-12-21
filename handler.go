package main

import (
	"bufio"
	"io"
	"net"
)

//Handler is a loop dispatcher
type Handler struct {
	conn      net.Conn
	processor func(input string) string
}

// Handles the incoming connection
func (handler *Handler) handle() {

	defer handler.conn.Close()

	// Read input from connection, keep it open until client exits
	scanner := bufio.NewScanner(handler.conn)

	for scanner.Scan() {

		io.WriteString(handler.conn, handler.processor(scanner.Text()))
	}
}
