package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Product data type for export
type Product struct {
	ID          int
	Name        string
	Price       float32
	Description string
}

var tpl *template.Template
var db *sql.DB

func main() {
	// Parse templates and handle error
	var err error
	tpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		fmt.Println("Error parsing templates:", err)
		return
	}

	// Open database connection
	db, err = sql.Open("mysql", "root:duynghia123@tcp(localhost:3306)/testdb")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// Verify connection
	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	fmt.Println("Database connection successful")

	// Set up routes
	http.HandleFunc("/productsearch", productSearchHandler)
	http.HandleFunc("/", homePageHandler)

	// Start server
	fmt.Println("Starting server on localhost:8080")
	err = http.ListenAndServe(":8080", nil) // Changed to listen on all interfaces
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func productSearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "productsearch.html", nil)
		return
	}

	// Handle POST request
	r.ParseForm()
	name := r.FormValue("productName")
	fmt.Println("Searching for product:", name)

	// Query the database
	stmt := "SELECT id, name, price, description FROM products WHERE name = ?"
	var p Product
	err := db.QueryRow(stmt, name).Scan(&p.ID, &p.Name, &p.Price, &p.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			// No product found
			tpl.ExecuteTemplate(w, "productsearch.html", struct {
				Error string
			}{"No product found with name: " + name})
			return
		}
		// Other errors (e.g., database issue)
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Product found, render template with data
	tpl.ExecuteTemplate(w, "productsearch.html", p)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}
