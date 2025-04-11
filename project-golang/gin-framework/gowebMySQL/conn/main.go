package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Hardcode the DSN with the password "duynghia123"
	dsn := "root:duynghia123@tcp(localhost:3306)/testdb"
	fmt.Println("Connecting with DSN:", dsn) // For debugging

	// Open database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error validating sql.Open arguments:", err)
		os.Exit(1)
	}
	defer db.Close()

	// Verify connection
	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		os.Exit(1)
	}

	// Insert data using Exec
	_, err = db.Exec("INSERT INTO users (firstname, lastname) VALUES ('Carl', 'Jones')")
	if err != nil {
		fmt.Println("Error executing INSERT statement:", err)
		os.Exit(1)
	}

	fmt.Println("Successful Connection and Data Insertion!")

	// Set up HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello! The app is running and connected to the database.")
	})

	// Define the port for the app
	port := "8080"
	fmt.Println("Starting server on port", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
