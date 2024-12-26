package main

import "fmt"

// Function add takes two integer parameters and returns their sum
func add(a int, b int) int {
	return a + b
}

func calculate(x, y int, a, b float64, name string) {
	fmt.Printf("x: %d, y: %d, a: %.2f, b: %.2f, name: %s\n", x, y, a, b, name)
}

func main() {
	sum := add(1, 2)
	calculate(1, 2, 3.14, 2.71, "Golang")
	fmt.Println("Sum is:", sum)
}
