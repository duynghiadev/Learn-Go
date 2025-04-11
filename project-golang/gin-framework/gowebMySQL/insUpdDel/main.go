package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Product data type
type Product struct {
	ID          int
	Name        string
	Price       float32
	Description string
}

var tpl *template.Template
var db *sql.DB

func main() {
	var err error
	// Parse templates with error handling
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
	if err = db.Ping(); err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	fmt.Println("Database connection successful")

	// Set up routes
	http.HandleFunc("/insert", insertHandler)
	http.HandleFunc("/browse", browseHandler)
	http.HandleFunc("/update", updateHandler) // Removed trailing slash
	http.HandleFunc("/updateresult", updateResultHandler)
	http.HandleFunc("/delete", deleteHandler) // Removed trailing slash
	http.HandleFunc("/", homePageHandler)

	// Start server
	fmt.Println("Starting server on :8080")
	if err = http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func browseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****browseHandler running*****")
	rows, err := db.Query("SELECT id, name, price, description FROM products")
	if err != nil {
		http.Error(w, "Error querying products: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.Description)
		if err != nil {
			http.Error(w, "Error scanning products: "+err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, p)
	}
	tpl.ExecuteTemplate(w, "select.html", products)
}

func insertHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****insertHandler running*****")
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "insert.html", nil)
		return
	}

	r.ParseForm()
	name := r.FormValue("nameName")
	priceStr := r.FormValue("priceName")
	descr := r.FormValue("descrName")

	if name == "" || priceStr == "" || descr == "" {
		tpl.ExecuteTemplate(w, "insert.html", "Error: All fields are required")
		return
	}

	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		tpl.ExecuteTemplate(w, "insert.html", "Error: Invalid price format")
		return
	}

	ins, err := db.Prepare("INSERT INTO products (name, price, description) VALUES (?, ?, ?)")
	if err != nil {
		http.Error(w, "Error preparing insert: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer ins.Close()

	res, err := ins.Exec(name, price, descr)
	if err != nil {
		tpl.ExecuteTemplate(w, "insert.html", "Error inserting data: "+err.Error())
		return
	}

	rowsAff, err := res.RowsAffected()
	if err != nil || rowsAff != 1 {
		tpl.ExecuteTemplate(w, "insert.html", "Error: No rows affected")
		return
	}

	lastID, _ := res.LastInsertId()
	fmt.Printf("Inserted ID: %d, Rows affected: %d\n", lastID, rowsAff)
	tpl.ExecuteTemplate(w, "insert.html", "Product Successfully Inserted")
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****updateHandler running*****")
	idStr := r.URL.Query().Get("idproducts")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Redirect(w, r, "/browse", http.StatusTemporaryRedirect)
		return
	}

	var p Product
	err = db.QueryRow("SELECT id, name, price, description FROM products WHERE id = ?", id).
		Scan(&p.ID, &p.Name, &p.Price, &p.Description)
	if err != nil {
		fmt.Println("Error fetching product:", err)
		http.Redirect(w, r, "/browse", http.StatusTemporaryRedirect)
		return
	}
	tpl.ExecuteTemplate(w, "update.html", p)
}

func updateResultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****updateResultHandler running*****")
	if r.Method != "POST" {
		http.Redirect(w, r, "/browse", http.StatusSeeOther)
		return
	}

	r.ParseForm()
	idStr := r.URL.Query().Get("idproducts")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		tpl.ExecuteTemplate(w, "result.html", "Error: Invalid product ID")
		return
	}

	name := strings.TrimSpace(r.FormValue("nameName"))
	priceStr := strings.TrimSpace(r.FormValue("priceName"))
	description := strings.TrimSpace(r.FormValue("descrName"))

	fmt.Println("Form values - ID:", id, "Name:", name, "Price:", priceStr, "Description:", description)

	if name == "" || priceStr == "" || description == "" {
		tpl.ExecuteTemplate(w, "result.html", "Error: All fields are required")
		return
	}

	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		fmt.Println("Error parsing price:", err, "Received value:", priceStr)
		tpl.ExecuteTemplate(w, "result.html", "Error: Invalid price format. Please enter a valid number (e.g., 999.99)")
		return
	}

	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ?, description = ? WHERE id = ?")
	if err != nil {
		http.Error(w, "Error preparing update: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(name, price, description, id)
	if err != nil {
		tpl.ExecuteTemplate(w, "result.html", "Error updating product: "+err.Error())
		return
	}

	rowsAff, err := res.RowsAffected()
	if err != nil {
		tpl.ExecuteTemplate(w, "result.html", "Error checking update: "+err.Error())
		return
	}
	if rowsAff == 0 {
		tpl.ExecuteTemplate(w, "result.html", "Error: No product found with ID "+idStr)
		return
	}

	tpl.ExecuteTemplate(w, "result.html", "Product was Successfully Updated")
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****deleteHandler running*****")
	idStr := r.URL.Query().Get("idproducts")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		tpl.ExecuteTemplate(w, "result.html", "Error: Invalid ID")
		return
	}

	del, err := db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		http.Error(w, "Error preparing delete: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer del.Close()

	res, err := del.Exec(id)
	if err != nil {
		tpl.ExecuteTemplate(w, "result.html", "Error deleting product: "+err.Error())
		return
	}

	rowsAff, err := res.RowsAffected()
	if err != nil || rowsAff != 1 {
		tpl.ExecuteTemplate(w, "result.html", "Error: No rows deleted")
		return
	}

	tpl.ExecuteTemplate(w, "result.html", "Product was Successfully Deleted")
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/browse", http.StatusTemporaryRedirect)
}
