package goroutine

import (
	"fmt"
	"sync"
	"time"
)

var Wg1 sync.WaitGroup // Fix: Separate WaitGroup for goroutine package

func CreatePizza1(pizza int) {
	defer Wg1.Done() // Fix: Ensure WaitGroup Done is called correctly
	time.Sleep(time.Second)
	fmt.Printf("Created Pizza goroutine %d \n", pizza)
}

func TimeTrack1(start time.Time, functionName string) {
	elapsed := time.Since(start)
	fmt.Println(functionName, "took", elapsed)
}
