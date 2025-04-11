# Gopher API using Gin Framework

## Overview

This Go application provides a simple REST API using the Gin framework to manage a list of Gophers.

## Code

```go
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

func getGopher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonGopher, err := json.Marshal(gophers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(jsonGopher)
}

func createGopher(w http.ResponseWriter, r *http.Request) {
	var newGopher gopher
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newGopher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	gophers = append(gophers, newGopher)
	GophersJSON, _ := json.Marshal(gophers)
	w.Header().Set("Content-Type", "application/json")
	w.Write(GophersJSON)
}

```

## API Endpoints

### 1. Get all Gophers

**Request:**

```sh
curl -X GET http://localhost:8080/gopher
```

**Response:**

```json
[
  { "id": "1", "first_name": "Ken", "last_name": "Thompson" },
  { "id": "2", "first_name": "Robert", "last_name": "Griesemer" }
]
```

### 2. Add a New Gopher

**Request:**

```sh
curl -X POST http://localhost:8080/gopher \
  -H "Content-Type: application/json" \
  -d '{"id": "3", "first_name": "Rob", "last_name": "Pike"}'
```

**Response:**

```json
[
  { "id": "1", "first_name": "Ken", "last_name": "Thompson" },
  { "id": "2", "first_name": "Robert", "last_name": "Griesemer" },
  { "id": "3", "first_name": "Rob", "last_name": "Pike" }
]
```

### 3. Bad Request Example

**Request:**

```sh
curl -X POST http://localhost:8080/gopher \
  -H "Content-Type: application/json" \
  -d '{"id": 4, "first_name": "Brian", "last_name": "Kernighan"'
```

**Response (HTTP 400 Bad Request):**

- Because missing closing bracket (`}`) in JSON data

```sh
unexpected EOF
```

## Summary of API Calls

| HTTP Method | Endpoint  | Description       |
| ----------- | --------- | ----------------- |
| `GET`       | `/gopher` | Fetch all gophers |
| `POST`      | `/gopher` | Add a new gopher  |

---

### Notes:

- `c.IndentedJSON` automatically sets the `Content-Type` as `application/json`.
- `c.BindJSON(&newGopher)` binds the JSON payload to a struct.
- The API runs on `localhost:8080`.

🚀 Enjoy coding with Gin!
