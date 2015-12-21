package main

import (
	"log"
	"net"
)

// Server is a TCP connection manager
type Server struct {
	net, laddr string
}

// ListenAndDispatch starts connections for a Server
func (server *Server) ListenAndDispatch(processor func(input string) string) {

	// Listen for TCP conncetions on localhost port 9090
	lt, err := net.Listen(server.net, server.laddr)
	if err != nil {
		log.Fatalln(err)
	}

	// Close conenction when we exit
	defer lt.Close()

	// Continually listen for and accept connections
	for {

		// If a connection is available, accept it
		conn, err := lt.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		// Delagate to handler
		handler := Handler{conn: conn, processor: processor}
		handler.handle()
	}
}
