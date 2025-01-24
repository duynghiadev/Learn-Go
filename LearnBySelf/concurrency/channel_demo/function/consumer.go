package function

import (
	"fmt"
	"time"
)

// Consumer function receives data from the channel
func Consumer(ch <-chan int) {
	for value := range ch {
		fmt.Printf("[Consumer] Received: %d\n", value)
		time.Sleep(300 * time.Millisecond) // Simulate processing
	}
	fmt.Println("[Consumer] Finished receiving.")
}
