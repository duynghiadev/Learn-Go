package examples

import "fmt"

func receiveAndSend(c chan int) {
	fmt.Printf("Received: %d\n", <-c)
	fmt.Printf("Sending 2...\n")
	c <- 2
}

func Example2() {
	myChan := make(chan int)

	go receiveAndSend(myChan)
	myChan <- 1

	fmt.Printf("Value from receiveAndSend: %d\n", <-myChan)
}
