package main

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

// --- Interfaces and Structs ---
type ProductManager interface {
	AddProduct(Product) error
	GetProduct(id int) (Product, error)
	UpdateProduct(id int, newProduct Product) error
	DeleteProduct(id int) error
	ListProducts() []Product
}

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

// Implements Stringer for Product
func (p Product) String() string {
	return fmt.Sprintf("Product[ID=%d, Name=%q, Price=%.2f, Stock=%d]", p.ID, p.Name, p.Price, p.Stock)
}

// ProductStore implements ProductManager
type ProductStore struct {
	products map[int]Product
}

func (ps *ProductStore) AddProduct(p Product) error {
	if _, exists := ps.products[p.ID]; exists {
		return fmt.Errorf("product with ID %d already exists", p.ID)
	}
	ps.products[p.ID] = p
	return nil
}

func (ps *ProductStore) GetProduct(id int) (Product, error) {
	if p, exists := ps.products[id]; exists {
		return p, nil
	}
	return Product{}, fmt.Errorf("product with ID %d not found", id)
}

func (ps *ProductStore) UpdateProduct(id int, newProduct Product) error {
	if _, exists := ps.products[id]; !exists {
		return fmt.Errorf("cannot update product with ID %d: not found", id)
	}
	ps.products[id] = newProduct
	return nil
}

func (ps *ProductStore) DeleteProduct(id int) error {
	if _, exists := ps.products[id]; !exists {
		return fmt.Errorf("cannot delete product with ID %d: not found", id)
	}
	delete(ps.products, id)
	return nil
}

func (ps *ProductStore) ListProducts() []Product {
	var list []Product
	for _, p := range ps.products {
		list = append(list, p)
	}
	return list
}

// --- Utility Functions ---
func printProductDetails(i interface{}) {
	switch v := i.(type) {
	case Product:
		fmt.Printf("Product Details: %s\n", v)
	case error:
		fmt.Printf("Error: %v\n", v)
	default:
		fmt.Printf("Unknown Type: %T\n", v)
	}
}

// --- Example with Readers ---
func readProductFromInput(input io.Reader) (Product, error) {
	var id int
	var name string
	var price float64
	var stock int

	_, err := fmt.Fscanf(input, "%d %s %f %d", &id, &name, &price, &stock)
	if err != nil {
		return Product{}, errors.New("invalid input format, expected: ID Name Price Stock")
	}

	return Product{ID: id, Name: name, Price: price, Stock: stock}, nil
}

// --- Main Function ---
func main() {
	store := &ProductStore{products: make(map[int]Product)}

	// Adding products
	fmt.Println("Adding Products:")
	store.AddProduct(Product{ID: 1, Name: "Laptop", Price: 999.99, Stock: 10})
	store.AddProduct(Product{ID: 2, Name: "Mouse", Price: 25.50, Stock: 50})
	store.AddProduct(Product{ID: 3, Name: "Keyboard", Price: 45.00, Stock: 30})

	// Listing products
	fmt.Println("\nListing Products:")
	for _, p := range store.ListProducts() {
		printProductDetails(p)
	}

	// Getting a product
	fmt.Println("\nGetting Product with ID 2:")
	product, err := store.GetProduct(2)
	printProductDetails(product)
	printProductDetails(err)

	// Updating a product
	fmt.Println("\nUpdating Product with ID 3:")
	err = store.UpdateProduct(3, Product{ID: 3, Name: "Mechanical Keyboard", Price: 75.00, Stock: 20})
	printProductDetails(err)

	// Listing products after update
	fmt.Println("\nListing Products After Update:")
	for _, p := range store.ListProducts() {
		printProductDetails(p)
	}

	// Deleting a product
	fmt.Println("\nDeleting Product with ID 1:")
	err = store.DeleteProduct(1)
	printProductDetails(err)

	// Listing products after deletion
	fmt.Println("\nListing Products After Deletion:")
	for _, p := range store.ListProducts() {
		printProductDetails(p)
	}

	// Reading a product from input
	fmt.Println("\nReading Product From Input:")
	input := strings.NewReader("4 Monitor 199.99 15")
	newProduct, err := readProductFromInput(input)
	printProductDetails(err)
	printProductDetails(newProduct)

	// Adding the new product
	fmt.Println("\nAdding New Product From Input:")
	err = store.AddProduct(newProduct)
	printProductDetails(err)

	// Final list of products
	fmt.Println("\nFinal Product List:")
	for _, p := range store.ListProducts() {
		printProductDetails(p)
	}
}
