package main

import "fmt"

// FindFirst returns the first occurrence of x in slice s
// Returns -1 if not found
func FindFirst[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Example with integers
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Index of 3 in numbers: %d\n", FindFirst(numbers, 3))
	fmt.Printf("Index of 6 in numbers: %d\n", FindFirst(numbers, 6))

	// Example with strings
	words := []string{"apple", "banana", "orange"}
	fmt.Printf("Index of 'banana' in words: %d\n", FindFirst(words, "banana"))
	fmt.Printf("Index of 'grape' in words: %d\n", FindFirst(words, "grape"))

	// Example with custom type
	type Person struct {
		name string
		age  int
	}
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}
	fmt.Println("Person:", people)
	searchPerson := Person{"Bob", 30}
	fmt.Printf("Index of %v in people: %d\n", searchPerson, FindFirst(people, searchPerson))
}
