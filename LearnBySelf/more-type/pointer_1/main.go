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
	p := Person{Name: "Alice", Age: 23}
	fmt.Println(p)
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

	grade, exists := grades["David"]
	if exists {
		fmt.Println("David's grade:", grade)
	} else {
		fmt.Println("David not found")
	}

	// --- Pointers Example ---
	fmt.Println("\n--- Pointers Example ---")
	// pointer to a struct
	personPointer := &p
	fmt.Println("Person struct pointer:", *personPointer)
	fmt.Println("Pointer to Person:", personPointer)
	fmt.Println("Name through pointer:", personPointer.Name)

	// Modify struct though pointer
	personPointer.Age = 30
	fmt.Println("Updated Age though pointer:", p.Age)
	// p start
	fmt.Println("p start:", p)

	// Pointer to a slice
	slicePointer := &numbers
	fmt.Println("Pointer to Slice:", slicePointer)
	fmt.Println("Slice through pointer:", *slicePointer)
	(*slicePointer)[0] = 10
	fmt.Println("Modified Slice through pointer:", numbers)

	// Pointer to a map
	mapPointer := &grades
	fmt.Println("Pointer to Map:", mapPointer)
	fmt.Println("Map through pointer:", *mapPointer)
	(*mapPointer)["Alice"] = 100
	fmt.Println("Modified Map through pointer:", grades)
}
