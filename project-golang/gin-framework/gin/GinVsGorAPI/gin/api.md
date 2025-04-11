# Gopher API using Gin Framework

## Overview

This Go application provides a simple REST API using the Gin framework to manage a list of Gophers.

## Code

```go
package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Gopher struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var gophers = []Gopher{
	{"1", "Ken", "Thompson"},
	{"2", "Robert", "Griesemer"},
}

func main() {
	router := gin.Default()
	router.GET("/gopher", getGophers)
	router.POST("/gopher", createGopher)
	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}

func getGophers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gophers)
}

func createGopher(c *gin.Context) {
	var newGopher Gopher
	err := c.BindJSON(&newGopher)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	gophers = append(gophers, newGopher)
	c.IndentedJSON(http.StatusCreated, gophers)
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
