package common

import "fmt"

// Pair struct to save name and value occurrences
type Pair struct {
	Name  string
	Count int
}

// Function to modify a slice
func ModifySlice(s []float64) {
	for i := range s {
		s[i] *= 1.1 // Increase each price by 10%
	}
}

// Function to update a map
func UpdateMap(m map[string]int, key string, newValue int) {
	if _, exists := m[key]; exists {
		m[key] = newValue
	}
}

// Function to update an array (will not modify the original array due to value semantics)
func UpdateArray(a [3]int) [3]int {
	for i := range a {
		a[i] += 10
	}
	return a
}

// CountOccurrences1 uses slice instead of map
func CountOccurrences1(names1 []string) {
	occurrences := []Pair{}

	for step, name := range names1 {
		// Log bước hiện tại
		fmt.Printf("\nStep %d: Processing name '%s'\n", step+1, name)

		// Check if the name already exists in the slice
		found := false
		for i, value := range occurrences {
			fmt.Printf("i: %d, Name: %s, Count: %d\n", i, value.Name, value.Count)
			if value.Name == name {
				occurrences[i].Count++
				found = true
				break
			}
		}
		// If the name doesn't exist, add a new entry
		if !found {
			occurrences = append(occurrences, Pair{Name: name, Count: 1})
		}

		// Log trạng thái hiện tại của object occurrences
		fmt.Println("Current state of occurrences:")
		for _, pair := range occurrences {
			fmt.Printf("  Name: %s, Count: %d\n", pair.Name, pair.Count)
		}
	}

	// Final log of all occurrences
	fmt.Println("\nFinal occurrences of each name:")
	for _, pair := range occurrences {
		fmt.Println("pair:", pair)
		fmt.Printf("%s: %d\n", pair.Name, pair.Count)
	}
}

// Function to count occurrences of names using a map
func CountOccurrences(names []string) {
	occurrences := make(map[string]int)
	for _, name := range names {
		occurrences[name]++
	}
	fmt.Println("\nOccurrences of each name:")
	for name, count := range occurrences {
		fmt.Printf("%s: %d\n", name, count)
	}
}

// Function to demonstrate slice manipulation
func PrintAndModifySlices(names []string) {
	// Creating a new slice
	newSlice := names[2:] // Slice from the third element onward
	fmt.Println("\nNew Slice from third element onward:", newSlice)

	// Modify the new slice
	newSlice[0] = "YYY"
	fmt.Println("\nAfter modifying the new slice:")
	fmt.Println("New Slice:", newSlice)
	fmt.Println("Original Array after modification via slice:", names)
}
