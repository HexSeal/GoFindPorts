package main

import (
	"fmt"
	"net"
)

func main() {
	// For testing, scan the first 1024 ports
	for i := 1; i <= 1024; i++ {
		// Get this iteration's address
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		
		// Establish the connection
		conn, err := net.Dial("tcp", address)

		// Handle errors
		if err != nil {
			// If we can't access the port, because it's either closed or filtered
			continue
		}

		// Close the connection once it's done to be good internet neighbors
		conn.Close()
		fmt.Printf("%d is open\n", i)
	}
}