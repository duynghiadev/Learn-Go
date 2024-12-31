package main

import (
	"fibonacci_closure/fibonacci"
	"fibonacci_closure/tracker"
	"fmt"
)

func main() {
	// Initialize Fibonacci tracker
	fibTracker := tracker.NewFibonacciTracker()
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
	standaloneFib := fibonacci.NewFibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", standaloneFib())
	}
	fmt.Println()
}
