package main

import "fmt"

// Define a struct for Product with embedded struct for Manufacturer
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
	// --- Advanced Structs Example ---
	fmt.Println("--- Advanced Structs Example ---")

	// Create a product with manufacturer details
	p := Product{
		Name:  "Laptop",
		Price: 1200.50,
		Stock: 15,
		Manufacturer: Manufacturer{
			Name:    "TechCorp",
			Country: "USA",
		},
	}

	fmt.Println("Product Details:")
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Price: $%.2f\n", p.Price)
	fmt.Printf("Stock: %d\n", p.Stock)
	fmt.Printf("Manufacturer: %s (%s)\n", p.Manufacturer.Name, p.Manufacturer.Country)

	// Modify struct values
	p.Stock -= 5 // Sell 5 units
	fmt.Println("\nUpdated Product Details:")
	fmt.Printf("Stock after sale: %d\n", p.Stock)

	// Use a pointer to modify nested struct
	manufacturerPointer := &p.Manufacturer
	manufacturerPointer.Country = "Canada"
	fmt.Printf("Manufacturer Country updated: %s\n", p.Manufacturer.Country)

	// Anonymous struct example
	fmt.Println("\n--- Anonymous Struct Example ---")
	anonymousProduct := struct {
		Name  string
		Price float64
	}{
		Name:  "Tablet",
		Price: 300.00,
	}
	fmt.Printf("Anonymous Product: %+v\n", anonymousProduct)
}
