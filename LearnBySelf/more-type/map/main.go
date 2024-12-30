package main

import "fmt"

// Vertex represents a geographical coordinate
type Vertex struct {
	Lat, Long float64
}

type Country struct {
	Name   string
	Cities []string
}

func main() {
	// Initialize a map with string keys and Vertex values
	locations := make(map[string]Vertex)

	// Add entries to the map
	locations["Bell Labs"] = Vertex{40.68433, -74.39967}
	locations["Google HQ"] = Vertex{37.42202, -122.08408}

	// Print the entire map
	fmt.Println("All locations:", locations)

	// Access and print a specific value
	fmt.Println("Bell Labs coordinates:", locations["Bell Labs"])

	// Update an entry
	locations["Google HQ"] = Vertex{37.422, -122.084}
	fmt.Println("Updated Google HQ coordinates:", locations["Google HQ"])

	// Check if a key exists
	key := "Microsoft HQ"
	if _, exists := locations[key]; exists {
		fmt.Println(key, "exists in the map.")
	} else {
		fmt.Println(key, "does not exist in the map.")
	}

	// Delete an entry
	delete(locations, "Bell Labs")
	fmt.Println("After deletion:", locations)

	// Iterate over the map
	fmt.Println("Iterating over locations:")
	for name, coords := range locations {
		fmt.Printf("%s: %+v\n", name, coords)
	}

	// Map of slices (advanced use case)
	countries := map[string][]string{
		"USA":   {"New York", "San Francisco", "Los Angeles"},
		"India": {"Delhi", "Mumbai", "Bangalore"},
	}

	// other type, but same as above (map of slices)
	countries1 := []Country{
		{Name: "USA", Cities: []string{"San Jose", "California", "Florida"}},
		{Name: "Vietnam", Cities: []string{"Hanoi", "Da Nang", "Ho Chi Minh"}},
	}

	fmt.Println("Cities by country map of slices:")
	for country, cities := range countries {
		fmt.Printf("%s: %v\n", country, cities)
	}

	fmt.Println("Cities by country slices of struct:")
	for _, country := range countries1 {
		fmt.Printf("%s: %v\n", country.Name, country.Cities)
	}
}
