package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Searchable interface defines methods for searching
type Searchable interface {
	comparable
	IsMatch(other any) bool
}

// FindFirst returns the first occurrence of x in slice s
func FindFirst[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

// FindAll returns all indices where x appears in slice s
func FindAll[T comparable](s []T, x T) []int {
	var indices []int
	for i, v := range s {
		if v == x {
			indices = append(indices, i)
		}
	}
	return indices
}

// FindBy returns the first element that matches the predicate
func FindBy[T any](s []T, predicate func(T) bool) (T, bool) {
	for _, v := range s {
		if predicate(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}

// Filter returns a new slice containing only elements that match the predicate
func Filter[T any](s []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range s {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Max returns the maximum element in a slice
func Max[T constraints.Ordered](s []T) (T, error) {
	if len(s) == 0 {
		var zero T
		return zero, fmt.Errorf("empty slice")
	}

	max := s[0]
	for _, v := range s[1:] {
		if v > max {
			max = v
		}
	}
	return max, nil
}

// Person struct with enhanced functionality
type Person struct {
	name string
	age  int
	id   string
}

// IsMatch implements Searchable interface
func (p Person) IsMatch(other any) bool {
	if otherPerson, ok := other.(Person); ok {
		return p.id == otherPerson.id
	}
	return false
}

func main() {
	// Basic type examples
	numbers := []int{1, 2, 3, 2, 4, 2, 5}
	fmt.Printf("First occurrence of 2: %d\n", FindFirst(numbers, 2))
	fmt.Printf("All occurrences of 2: %v\n", FindAll(numbers, 2))

	// Using FindBy with a predicate
	evenNumber, found := FindBy(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("First even number: %d, found: %v\n", evenNumber, found)

	// Using Filter
	evenNumbers := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("Even numbers: %v\n", evenNumbers)

	// Using Max
	max, err := Max(numbers)
	if err == nil {
		fmt.Printf("Maximum number: %d\n", max)
	}

	// Complex type example with Person
	people := []Person{
		{"Alice", 25, "001"},
		{"Bob", 30, "002"},
		{"Charlie", 35, "003"},
		{"David", 30, "004"},
	}

	// Find people by age
	thirtyYearOlds := Filter(people, func(p Person) bool {
		return p.age == 30
	})
	fmt.Printf("People who are 30: %+v\n", thirtyYearOlds)

	// Find person by custom predicate
	person, found := FindBy(people, func(p Person) bool {
		return p.name == "Alice"
	})
	if found {
		fmt.Printf("Found person: %+v\n", person)
	}
}
