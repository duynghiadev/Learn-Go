package main

import (
	"fmt"
	"goroutine_demo/number"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// Create two goroutines with proper synchronization
	go func() {
		defer wg.Done()
		number.PrintNumbers("Goroutine 1")
	}()

	go func() {
		defer wg.Done()
		number.PrintNumbers("Goroutine 2")
	}()

	fmt.Println("Main: Waiting for goroutines...")
	wg.Wait()
	fmt.Println("Main: Done!")
}
