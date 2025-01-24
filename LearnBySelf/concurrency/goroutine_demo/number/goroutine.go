package number

import (
	"fmt"
	"sync"
	"time"
)

// NumberPrinter represents a structure to handle number printing operations
type NumberPrinter struct {
	name    string
	done    chan bool
	numbers chan int
	wg      *sync.WaitGroup
}

// NewNumberPrinter creates a new NumberPrinter instance
func NewNumberPrinter(name string) *NumberPrinter {
	return &NumberPrinter{
		name:    name,
		done:    make(chan bool),
		numbers: make(chan int, 5), // buffered channel
		wg:      &sync.WaitGroup{},
	}
}

// PrintNumbers prints numbers with enhanced control and concurrent features
func PrintNumbers(name string) {
	printer := NewNumberPrinter(name)
	printer.wg.Add(2) // for producer and consumer goroutines

	// Start producer goroutine
	go func() {
		defer printer.wg.Done()
		for i := 1; i <= 5; i++ {
			select {
			case <-printer.done:
				return
			case printer.numbers <- i:
				time.Sleep(500 * time.Millisecond)
			}
		}
		close(printer.numbers)
	}()

	// Start consumer goroutine
	go func() {
		defer printer.wg.Done()
		for num := range printer.numbers {
			fmt.Printf("%s: %d\n", printer.name, num)
		}
	}()

	// Wait for completion or timeout
	done := make(chan bool)
	go func() {
		printer.wg.Wait()
		done <- true
	}()

	select {
	case <-done:
		fmt.Printf("%s: Completed successfully\n", name)
	case <-time.After(5 * time.Second):
		close(printer.done) // Signal to stop if timeout
		fmt.Printf("%s: Timed out\n", name)
	}
}
