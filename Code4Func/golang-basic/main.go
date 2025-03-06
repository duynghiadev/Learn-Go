package main

import (
	"fmt"
	"go-concurrency/concurrency/channel"
	"go-concurrency/concurrency/goroutine"
	"go-concurrency/concurrency/parallelism"
	"runtime"
	"time"
)

func main() {
	// start channel
	fmt.Println("---------------------channel---------------------")
	c := make(chan string)
	done := make(chan bool)

	go channel.Pinger(c)
	go channel.Printer(c, done)

	<-done // Wait until Printer signals completion

	fmt.Println("---------------------parallelism---------------------")

	// start parallelism
	parallelismDone := make(chan bool)
	start := time.Now()

	number_of_ovens := 3
	runtime.GOMAXPROCS(number_of_ovens)
	parallelism.Wg.Add(number_of_ovens)

	for i := 0; i < number_of_ovens; i++ {
		go parallelism.CreatePizza(i)
	}
	parallelism.Wg.Wait()

	go func() {
		parallelism.TimeTrack(start, "Build Pizzas", parallelismDone)
	}()
	<-parallelismDone // Wait for the time tracking to complete

	// end parallelism

	fmt.Println("---------------------goroutine---------------------")

	// start goroutine
	goroutineDone := make(chan bool)
	start = time.Now() // Reset start time for next section

	go func() {
		defer goroutine.TimeTrack1(start, "Build Pizzas1") // Track time
		goroutine.Wg1.Add(3)                               // Fix: Add to WaitGroup before spawning goroutines
		for i := 0; i < 3; i++ {
			go goroutine.CreatePizza1(i)
		}
		goroutine.Wg1.Wait() // Fix: Wait for all goroutines to complete
		goroutineDone <- true
	}()

	<-goroutineDone // Ensure goroutine section completes before exiting

	// end goroutine
}
