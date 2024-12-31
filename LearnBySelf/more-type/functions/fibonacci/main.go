package main

import (
	"fmt"
)

// FibonacciTracker tracks Fibonacci numbers and their frequencies.
type FibonacciTracker struct {
	cache map[int]int // Cache to store precomputed Fibonacci numbers
	calls map[int]int // Tracks how many times each Fibonacci number was returned
}

// NewFibonacciTracker initializes a new FibonacciTracker.
func NewFibonacciTracker() *FibonacciTracker {
	return &FibonacciTracker{
		cache: make(map[int]int),
		calls: make(map[int]int),
	}
}

// FibonacciClosure generates a closure for successive Fibonacci numbers.
func (ft *FibonacciTracker) FibonacciClosure() func() int {
	n1, n2 := 0, 1
	index := 0

	return func() int {
		var result int
		if val, exists := ft.cache[index]; exists {
			result = val
		} else {
			if index == 0 {
				result = n1
			} else if index == 1 {
				result = n2
			} else {
				result = n1 + n2
				n1, n2 = n2, result
			}
			ft.cache[index] = result
		}
		ft.calls[result]++
		index++
		return result
	}
}

// PrintHistory prints the Fibonacci numbers and their frequencies.
func (ft *FibonacciTracker) PrintHistory() {
	fmt.Println("\nFibonacci History:")
	for num, count := range ft.calls {
		fmt.Printf("Number: %d, Count: %d\n", num, count)
	}
}

// NewFibonacci generates a standalone Fibonacci closure.
func NewFibonacci() func() int {
	n1, n2 := 0, 1
	return func() int {
		result := n1
		n1, n2 = n2, n1+n2
		return result
	}
}

func main() {
	// Initialize Fibonacci tracker
	fibTracker := NewFibonacciTracker()
	fibonacciClosure := fibTracker.FibonacciClosure()

	// Generate Fibonacci numbers
	fmt.Println("Generated Fibonacci Numbers:")
	for i := 0; i < 20; i++ {
		fmt.Printf("%d ", fibonacciClosure())
	}

	// Print Fibonacci history
	fibTracker.PrintHistory()

	// Demonstrate standalone Fibonacci generator
	fmt.Println("\nStandalone Fibonacci Generator:")
	standaloneFib := NewFibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", standaloneFib())
	}
	fmt.Println()
}
