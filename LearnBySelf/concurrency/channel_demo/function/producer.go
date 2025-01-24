package function

import (
	"fmt"
	"time"
)

// Producer function sends data to the channel
func Producer(ch chan<- int, count int) {
	for i := 1; i <= count; i++ {
		fmt.Printf("[Producer] Sending: %d\n", i)
		ch <- i
		time.Sleep(500 * time.Millisecond) // Simulate work
	}
	close(ch) // Close the channel after sending all data
	fmt.Println("[Producer] Finished sending.")
}
