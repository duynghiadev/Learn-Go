package main

import "fmt"

// Define a struct for Product
type Product struct {
	Name  string
	Price int
	Stock int
}

func main() {
	// --- Structs Example ---
	fmt.Println("--- Structs Example ---")
	p := Product{Name: "Laptop", Price: 50000, Stock: 10}
	fmt.Printf("Product: %+v\n", p)
	fmt.Println("Name:", p.Name)
	fmt.Println("Price:", p.Price)
	fmt.Println("Stock:", p.Stock)
}
