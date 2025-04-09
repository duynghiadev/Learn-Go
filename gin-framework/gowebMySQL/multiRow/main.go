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
	tpl, _ = template.ParseGlob("templates/*.html")
	var err error
	// Updated connection string with password "duynghia123"
	db, err = sql.Open("mysql", "root:duynghia123@tcp(localhost:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err.Error())
	}
	fmt.Println("Database connection successful")

	http.HandleFunc("/productsearch", productSearchHandler)
	http.HandleFunc("/productsearch2", productSearchHandler2)
	http.HandleFunc("/", homePageHandler)
	http.ListenAndServe("localhost:8080", nil)
}

func productSearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "productsearch.html", nil)
		return
	}

	r.ParseForm()
	min := r.FormValue("minPriceName")
	max := r.FormValue("maxPriceName")

	stmt := "SELECT * FROM products WHERE price >= ? && price <= ?;"
	rows, err := db.Query(stmt, min, max)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.Description)
		if err != nil {
			panic(err)
		}
		products = append(products, p)
	}
	tpl.ExecuteTemplate(w, "productsearch.html", products)
}

func productSearchHandler2(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "productsearch2.html", nil)
		return
	}

	r.ParseForm()
	min := r.FormValue("minPriceName")
	max := r.FormValue("maxPriceName")
	stmt, err := db.Prepare("SELECT * FROM products WHERE price >= ? && price <= ?;")

	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(min, max)
	var products []Product

	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.Description)
		if err != nil {
			panic(err)
		}
		products = append(products, p)
	}
	tpl.ExecuteTemplate(w, "productsearch2.html", products)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}
