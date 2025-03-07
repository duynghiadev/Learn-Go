package main

import (
	"fmt"
	"runtime"
	"sync"
)

func g1() {
	fmt.Println("g1()")
}

func g2() {
	fmt.Println("g2()")
}

func g3() {
	fmt.Println("g3()")
}

func g4() {
	fmt.Println("g4()")
}

func main() {
	// counter CPU have in local laptop
	numberP := runtime.NumCPU()
	fmt.Println("numberP:", numberP)

	fmt.Println("-------------------")

	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(4) // Add counter for 4 goroutines

	go func() {
		g1()
		wg.Done()
	}()

	go func() {
		g2()
		wg.Done()
	}()

	go func() {
		g3()
		wg.Done()
	}()

	go func() {
		g4()
		wg.Done()
	}()

	wg.Wait() // Wait for all goroutines to complete
}
