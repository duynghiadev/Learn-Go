package main

import (
	"channel_golang/function"
	"time"
)

func main() {
	// Create a channel for communication
	ch := make(chan int)

	// Number of items to produce
	itemCount := 5

	// Start producer and consumer as goroutines
	go function.Producer(ch, itemCount)
	go function.Consumer(ch)

	// Wait for producer and consumer to finish
	time.Sleep(4 * time.Second)

	// Call AddProduct to demonstrate the Sum function
	function.AddProduct()

	println("Program completed.")
}
