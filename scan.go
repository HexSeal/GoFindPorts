package main

import (
	"fmt"
	"net"
)

func main() {
	// First query is type, second is where to scan(testing site)
	_, err := net.Dial("tcp", "scanmen.nmap.org:80")
	
	// If there is no error, print telling the user the connection was successful
	if err == nil {
		fmt.Println("Connection successful")
	}
}