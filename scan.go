package main

import (
	"fmt"
	"sync"
)

// Creates a worker pool with a wait group
func worker(ports chan int, wg *sync.WaitGroup) {
	// Keep receiving work from the ports channel
	for port := range ports {
		fmt.Println(port)
		wg.Done()
	}
}

func main() {
	// Create a channel, and give it a buffer of 100 items. This allows workers to start immediately
	ports := make(chan int, 100)
	// Create the wait group
	var wg sync.WaitGroup

	// Start up the workers
	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
	}

	// Send a port to a worker in the ports channel
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i
	}
	// Wait until the program is finished, then close all open connections
	wg.Wait()
	close(ports)
}
