package channel

import (
	"fmt"
	"time"
)

func Pinger(c chan string) {
	for i := 0; i < 3; i++ {
		c <- "ping"
	}
	close(c) // Close the channel after sending all pings
}

func Printer(c chan string, done chan bool) {
	count := 0
	for msg := range c { // This will automatically exit when channel is closed
		fmt.Printf("Message %d: %s\n", count+1, msg)
		time.Sleep(time.Second * 1)
		count++
	}
	fmt.Printf("Receive total of %d messages \n", count)
	done <- true // Signal that we're done
}
