package main

import (
	"fmt"
	"math"
	"my-app/operations"
	"my-app/tracker"
	"strconv"
)

func main() {
	// Define operations
	ops := []operations.Operation{
		{Name: "Addition", Action: func(a, b float64) float64 { return a + b }},
		{Name: "Subtraction", Action: func(a, b float64) float64 { return a - b }},
		{Name: "Multiplication", Action: func(a, b float64) float64 { return a * b }},
		{Name: "Division", Action: func(a, b float64) float64 {
			if b != 0 {
				return a / b
			}
			return math.NaN() // Return NaN for division by zero
		}},
		{Name: "Power", Action: math.Pow}, // Using math.Pow as a function
	}

	// Initialize a Tracker to maintain operation history
	tr := tracker.NewTracker()

	fmt.Println("Available operations:")
	for i, op := range ops {
		fmt.Printf("%d: %s\n", i+1, op.Name)
	}

	for {
		// Ask the user for input
		fmt.Print("\nEnter operation number (or 0 to quit): ")
		var input string
		fmt.Scanln(&input)

		// Convert input to integer
		choice, err := strconv.Atoi(input)
		if err != nil || choice < 0 || choice > len(ops) {
			fmt.Println("Invalid choice. Please try again.")
			continue
		}

		if choice == 0 {
			fmt.Println("Exiting program.")
			break
		}

		// Get the operation based on user choice
		op := ops[choice-1]

		// Ask for two numbers
		var x, y float64
		fmt.Print("Enter first number: ")
		fmt.Scanln(&x)
		fmt.Print("Enter second number: ")
		fmt.Scanln(&y)

		// Perform the operation
		result := tr.PerformOperation(op, x, y)
		fmt.Printf("Operation: %s, Inputs: (%.2f, %.2f), Result: %.2f\n", op.Name, x, y, result)
	}

	// Print operation history
	fmt.Println("\nOperation History:")
	for name, results := range tr.History {
		fmt.Println("tracker:", tr)
		fmt.Println("tracker.History:", tr.History)
		fmt.Println("name:", name, "results:", results)
		fmt.Println("----------------------")
		fmt.Printf("%s: %v\n", name, results)
	}
}
