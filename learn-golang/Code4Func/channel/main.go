package main

import (
	"fmt"
	"sync"
)

/*
   channel in golang:
   1. unbuffered channel
   2. buffered channel
   3. select
   4. close channel
*/

func unbufferedChannel() {
	ch := make(chan int)

	go func() {
		ch <- 100
		fmt.Println("Unbuffered channel: sent")
	}()

	fmt.Println("Unbuffered channel: received", <-ch)
	fmt.Println("Unbuffered channel: Done")
	fmt.Println("---------------------------")
}

func bufferedChannel() {
	ch2 := make(chan int, 2)

	ch2 <- 1
	ch2 <- 2

	close(ch2)

	fmt.Println("Buffered channel: received", <-ch2)
	fmt.Println("Buffered channel: received", <-ch2)

	// Check if the channel is closed and avoid reading from it after close.
	_, ok := <-ch2
	if !ok {
		fmt.Println("Buffered channel: Channel closed")
		fmt.Println("---------------------------")
	}
}

func selectExample() {
	queue := make(chan int)
	done := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(1) // Add a wait group.

	go func() {
		defer wg.Done() // Signal that the goroutine is done.
		for i := 0; i < 5; i++ {
			queue <- i
		}
		close(queue) // Close the queue channel when done.
	}()

	go func() {
		wg.Wait()   // Wait for the queue goroutine to finish.
		close(done) // Close the done channel when the queue is finished.
	}()

	for {
		select {
		case v, ok := <-queue:
			if !ok {
				// Channel is closed, no more values.
				queue = nil // Prevent further reads from the closed channel.
				continue
			}
			fmt.Println("Select: value:", v)
		case <-done:
			fmt.Println("Select: Done")
			return
		}
	}
}

func main() {
	// Unbuffered channel
	unbufferedChannel()

	// Buffered channel
	bufferedChannel()

	// Select with graceful shutdown
	selectExample()
}
