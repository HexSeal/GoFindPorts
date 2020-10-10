package main

import (
	"fmt"
	"net"
	"sync"
)

// Scans concurrently, but produces inconsistent results due to network limitations
func main() {
	// Struct instance created by the sync package
	var wg sync.WaitGroup
	// For testing, scan the first 1024 ports
	for i := 1; i <= 1024; i++ {
		// Add time to the wait group
		wg.Add(1)
		// Want to do it faster? Wrap it in a go routine.
		go func(j int) {
			// Defer subtracting from the wait group until we're done
			defer wg.Done()
			// Get this iteration's address
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			
			// Establish the connection
			conn, err := net.Dial("tcp", address)

			// Handle errors
			if err != nil {
				// If we can't access the port, because it's either closed or filtered
				return
			}

			// Close the connection once it's done to be good internet neighbors
			conn.Close()
			fmt.Printf("%d is open\n", j)
		}(i) 
	}
	// Stops the wait group froming being done until the function exits
	wg.Wait()
}