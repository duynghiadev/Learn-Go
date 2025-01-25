package main

import "fmt"

func receiveAndSend(c chan int) {
	fmt.Printf("Received: %d\n", <-c)
	fmt.Printf("Sending 2...\n")
	c <- 2
}

func receiveOnly(c <-chan int) {
	fmt.Printf("Received only chan: %d\n", c)
	// c <- 2 // error send
}

func sendOnly(c chan<- int) {
	c <- 2 // OK
	fmt.Printf("Send only chan: %d\n", c)
	// <-c /// error receive
}

func main() {
	myChan := make(chan int)

	go receiveAndSend(myChan)
	myChan <- 3
	fmt.Printf("Value from receiveAndSend: %d\n", <-myChan)

	go func() {
		for i := 1; i <= 5; i++ {
			myChan <- i
		}
		close(myChan)
	}()

	for value := range myChan {
		fmt.Printf("Value in for-range: %d\n", value)
	}

	for {
		value, isAlive := <-myChan

		if !isAlive {
			fmt.Printf("Value: %d. Channel has been closed\n", value)
			break
		}
		fmt.Printf("Value: %d\n", value)
	}
}
