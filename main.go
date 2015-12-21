package main

import (
	"strings"
)

// Holds the cache storage
var storage = make(map[string]string)

// Main entry point of application
func main() {

	// Creates server at given location
	server := Server{net: "tcp", laddr: ":9090"}

	// Set process() func defined here as the handler's processor
	handler := ScanHandler{processor: process}

	// Starts connection listening loop which it processes with given handler
	server.ListenAndDispatch(handler)
}

// Handles the incoming connection, gets input, returns output
func process(input string) string {

	// Initialize output to return char
	output := "\n"

	// Splits input string by spaces into an array
	parts := strings.Fields(input)

	// Check that we have atleast 2 parts in the input
	if len(parts) < 2 {
		return output
	}

	// Verb comes first: GET, SET, DEL
	verb := parts[0]

	// Always have a data key for all operations
	key := parts[1]

	// Depending on verb take specific action
	switch verb {
	case "GET":
		output = storage[key] + "\n"
	case "SET":
		storage[key] = parts[2]
	case "DEL":
		delete(storage, key)
	default:
		output = "Invalid command\n"
	}

	return output
}
