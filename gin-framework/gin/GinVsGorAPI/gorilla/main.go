package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type gopher struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var gophers = []gopher{
	{"1", "Ken", "Thompson"},
	{"2", "Robert", "Griesemer"},
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/gopher", getGopher).Methods("GET")
	router.HandleFunc("/gopher", createGopher).Methods("POST")
	http.Handle("/", router)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

/*
w.Header.Set() sets the content type to application/json
if w.WriteHeader is no called explicitly, implicit http.StatusOK is used with w.Write()
jsonGopher, err := json.Marshal(gopher) converts the gopher struct to JSON
if err != nil {} sends the 400 error code if the marshaling fails and ends the func
w.Writer()  writes the data to the connection as part of an HTTP reply
*/
func getGopher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonGopher, err := json.Marshal(gophers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(jsonGopher)
}

/*
A Decoder reads and decodes JSON values from an input stream.
NewDecoder returns a new decoder that reads from r.
The decoder introduces its own buffering and may read data from r beyond the JSON values requested.
Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by v.
handle any errors create decoder, setting HTTP status code
since err from defer r.Body.Close() is not nill, closing r.Body must be explicitly closed
json.Marshal(gophers) marshals gophers slice so w.Write() can receive it
w.Header().Set() sets the headers to application/json
w.Write() writes our JSON to the http.ResponseWriter
*/
func createGopher(w http.ResponseWriter, r *http.Request) {
	var newGopher gopher

	// Step 1: Decode JSON request body into newGopher struct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newGopher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // Function exits here if there's an error
	}

	// Step 2: Defer closing request body (executed at the END of function)
	defer r.Body.Close()

	// Step 3: Add newGopher to the gophers slice
	gophers = append(gophers, newGopher)

	// Step 4: Marshal the updated gophers list to JSON
	GophersJSON, _ := json.Marshal(gophers)

	// Step 5: Set response header and write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(GophersJSON)

	// Step 6: Function exits -> `defer r.Body.Close()` is executed here
}
