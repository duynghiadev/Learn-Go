package main

import (
	"fmt"
	"learn-by-self/more-type/slice/common"
)

type Vertex struct {
	X int
	Y int
}

type Manufacturer struct {
	Name    string
	Country string
}

type Product struct {
	Name         string
	Price        float64
	Stock        int
	Manufacturer Manufacturer
}

func main() {
	// --- Structs Field Example ---
	fmt.Println("\n--- Structs Field Example ---")
	v := Vertex{1, 2}
	fmt.Println("v has X before change = ", v)
	v.X = 4
	fmt.Println("v has X after change = ", v)
	fmt.Println("v.X = ", v.X)

	// --- Structs Field Pointer Example ---
	fmt.Println("\n--- Structs Field Pointer Example ---")
	p1 := &v
	p1.X = 1e9
	fmt.Println("v has X after change via pointer = ", v)

	// --- Structs Field Pointer Example 2 ---
	fmt.Println("\n--- Structs Field Pointer Example 2 ---")
	v1 := Vertex{1, 2}
	v2 := Vertex{X: 1} // Y:0 is implicit
	v3 := Vertex{}     // X:0 and Y:0
	p2 := &Vertex{3, 4}
	fmt.Println("v1 = ", v1)
	fmt.Println("v2 = ", v2)
	fmt.Println("v3 = ", v3)
	fmt.Println("p1 = ", p2)

	// --- Structs and Slice Example ---
	fmt.Println("\n--- Structs and Slice Example ---")
	v4 := []Vertex{
		{1, 2},
		{3, 4},
	}
	fmt.Println("v4 = ", v4)

	// --- Structs, Slice, Array, Map, Function Example ---
	fmt.Println("\n--- Structs, Slice, Array, Map, Function Example ---")
	var a [2]Vertex
	a[0] = Vertex{1, 2}
	a[1] = Vertex{3, 4}
	fmt.Println("a = ", a)

	// --- Structs, Slice, Array, Map, Function more advanced example ---
	fmt.Println("\n--- Structs, Slice, Array, Map, Function more advanced example ---")

	// --- Structs Example ---
	fmt.Println("\n--- Structs Example ---")
	p := Product{
		Name:  "Laptop",
		Price: 1000.00,
		Stock: 10,
		Manufacturer: Manufacturer{
			Name:    "Apple",
			Country: "USA",
		},
	}

	fmt.Println("Product Details:", p)
	fmt.Printf("Product Name: %s\n", p.Name)
	fmt.Printf("Product Price: %f\n", p.Price)
	fmt.Printf("Product Stock: %d\n", p.Stock)
	fmt.Printf("Product Manufacturer Name: %s (%s)\n", p.Manufacturer.Name, p.Manufacturer.Country)

	// Modify struct values
	p.Stock -= 5
	fmt.Println("Product Details after stock updated:", p)
	fmt.Println("Stock after sale:", p.Stock)

	// Use a pointer to modify nested struct
	manufacturePointer := &p.Manufacturer
	manufacturePointer.Country = "Singapore"
	fmt.Println("Product Details after manufacturer country updated:", p)

	// --- Slices Example ---
	fmt.Println("\n--- Slices Example ---")
	prices := []float64{19.99, 29.99, 39.99}
	fmt.Println("Initial Prices Slice:", prices)

	prices = append(prices, 49.99)
	fmt.Println("After Append:", prices)

	subPrices := prices[1:3]
	fmt.Println("Sub-slice of Prices:", subPrices)

	// Modify slice through a function
	common.ModifySlice(prices)
	fmt.Println("Prices after modification:", prices)

	// --- Arrays Example ---
	fmt.Println("\n--- Arrays Example ---")
	stocks := [3]int{100, 200, 300}
	fmt.Println("Initial Stocks Array:", stocks)

	common.UpdateArray(stocks)
	fmt.Println("Stocks after function call (no change):", stocks)

	// --- Maps Example ---
	fmt.Println("\n--- Maps Example ---")
	inventory := map[string]int{
		"Laptop": 15,
		"Tablet": 25,
		"Phone":  30,
	}

	fmt.Println("Initial Inventory Map:", inventory)

	inventory["Tablet"] -= 5 // Sell 5 tablets
	fmt.Println("Updated Inventory Map:", inventory)

	// Check and update using a function
	common.UpdateMap(inventory, "Laptop", 10)
	fmt.Println("Inventory after update function:", inventory)

	// --- Anonymous Struct Example ---
	fmt.Println("\n--- Anonymous Struct Example ---")
	anonymousProduct := struct {
		Name  string
		Price float64
	}{
		Name:  "Tablet",
		Price: 300.00,
	}
	fmt.Printf("Anonymous Product: %+v\n", anonymousProduct)

	fmt.Println("\n==============================")

	// Initialize an array of names
	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println("Original Array:", names)

	// Create slices from the array
	a1 := names[0:2]
	b1 := names[1:3]
	fmt.Println("Slice A:", a1)
	fmt.Println("Slice B:", b1)

	// Modify the shared underlying array through slice B
	b1[0] = "XXX"
	fmt.Println("\nAfter modifying Slice B:")
	fmt.Println("Slice A:", a1)
	fmt.Println("Slice B:", b1)
	fmt.Println("Modified Array:", names)

	// Enhanced Example: Using a map to count occurrences of names
	common.CountOccurrences(names[:]) // Pass the slice of the array

	// Using a function to print and manipulate slices
	common.PrintAndModifySlices(names[:])

	fmt.Printf("\n==============================\n")

	names1 := []string{"John", "Paul", "George", "John", "Paul", "John"}
	common.CountOccurrences1(names1)

}
