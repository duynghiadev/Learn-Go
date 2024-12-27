package main

import "fmt"

// Define a struct
type Person struct {
	Name string
	Age  int
}

func main() {
	// --- Structs Example ---
	fmt.Println("--- Structs Example ---")
	p := Person{Name: "Alice", Age: 25}
	fmt.Printf("%+v\n", p)
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)

	// --- Slices Example ---
	fmt.Println("\n--- Slices Example ---")
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", numbers)

	numbers = append(numbers, 6)
	fmt.Println("After append:", numbers)

	subSlice := numbers[1:4]
	fmt.Println("Sub-slice:", subSlice)

	// --- Maps Example ---
	fmt.Println("\n--- Maps Example ---")
	grades := map[string]int{
		"Alice": 90,
		"Bob":   85,
	}

	fmt.Println("Alice's grade:", grades["Alice"])

	grades["Charlie"] = 88
	fmt.Println("Updated map:", grades)

	grade, exists := grades["Bob"]

	fmt.Println("grade:", grade) // take the value of the key
	if exists {
		fmt.Println("David's grade:", grade)
	} else {
		fmt.Println("David not found")
	}
}
