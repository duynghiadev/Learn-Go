package main

import (
	"fmt"
	"golang-rest-api/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8081"

	router.HandleFunc("/", func(resp http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(resp, "Up and running...")
	})
	router.HandleFunc("/posts", handler.GetPosts).Methods("GET")

	log.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
