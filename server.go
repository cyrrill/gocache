package main

import (
	"log"
	"net"
)

// Server is a TCP connection manager
type Server struct {
	net, laddr string
}

// Handler interface defines what can be attached to server
type Handler interface {
	handle(conn net.Conn)
}

// ListenAndDispatch starts connections for a Server
func (server *Server) ListenAndDispatch(handler Handler) {

	// Listen for TCP conncetions on localhost port 9090
	lt, err := net.Listen(server.net, server.laddr)
	if err != nil {
		log.Fatalln(err)
	}

	// Close listener on exit
	defer lt.Close()

	// Continually listen for and accept connections
	for {

		// If a connection is available, accept it
		conn, err := lt.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		// Delegate to handler
		handler.handle(conn)
	}
}
