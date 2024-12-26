package main

import "fmt"

// Function add takes two integer parameters and returns their sum
func add(a int, b int) int {
	return a + b
}

func calculate(x, y int, a, b float64, name string) {
	fmt.Printf("Calculation: x: %d, y: %d, a: %.2f, b: %.2f, name: %s\n", x, y, a, b, name)
}

func swap(x, y string) (string, string) {
	return y, x
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// Example: Named Return Values
func divide2(a, b int) (result int, err error) {
	if b == 0 {
		err = fmt.Errorf("cannot divide by zero")
		return // Naked return: return 'result' and 'err' variables
	}
	result = a / b
	return // Naked return
}

func main() {
	sum := add(1, 2)

	calculate(1, 2, 3.14, 2.71, "Golang")

	fmt.Println("Sum is:", sum)

	first, second := swap("Hello", "World")

	fmt.Println("Swapped:", first, second)

	// divide
	result, error := divide(10, 2)
	if error != nil {
		fmt.Println("Error:", error)
	} else {
		fmt.Println("Result:", result)
	}

	// divide2
	result2, error2 := divide2(10, 5)
	if error2 != nil {
		fmt.Println("Error2:", error2)
	} else {
		fmt.Println("Result2:", result2)
	}
}
