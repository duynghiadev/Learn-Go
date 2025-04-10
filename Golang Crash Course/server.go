package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8081"

	router.HandleFunc("/", func(resp http.ResponseWriter, r *http.Request) {
		fmt.Println(resp, "Up and running...")
	})
	router.HandleFunc("/posts", getPosts).Methods("GET")
	log.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
