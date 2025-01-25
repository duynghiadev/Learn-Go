package examples

import (
	"fmt"
	"runtime"
)

func sender(c chan<- int, name string) {
	for i := 1; i <= 2; i++ {
		c <- 1
		fmt.Printf("%s has sent 1 to channel\n", name)
		runtime.Gosched()
	}
}

func Example9() {
	myChan := make(chan int)

	go sender(myChan, "S1")
	go sender(myChan, "S2")
	go sender(myChan, "S3")

	start := 0

	for {
		start += <-myChan
		fmt.Println(start)

		if start >= 6 {
			break
		}
	}
}
