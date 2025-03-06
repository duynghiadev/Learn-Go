package parallelism

import (
	"fmt"
	"sync"
	"time"
)

var Wg sync.WaitGroup

func CreatePizza(pizza int) {
	defer Wg.Done()
	time.Sleep(time.Second)
	fmt.Printf("Created Pizza parallelism %d \n", pizza)
}

func TimeTrack(start time.Time, functionName string, done chan bool) {
	elapsed := time.Since(start)
	fmt.Println(functionName, "took", elapsed)
	done <- true
}
