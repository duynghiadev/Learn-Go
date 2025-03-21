package examples

import "fmt"

func Example3() {
	myChan := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			myChan <- i
		}
		close(myChan)
	}()

	for {
		value, isAlive := <-myChan

		if !isAlive {
			fmt.Printf("Value: %d. Channel has been closed.\n", value)
			break
		}

		fmt.Printf("Value: %d\n", value)
	}
}
