package examples

import (
	"fmt"
	"time"
)

func Example1() {
	myChan := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			myChan <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 1; i <= 5; i++ {
		fmt.Println(<-myChan)
	}
}
