# Movie API using Gin Framework

## Overview

This Go application provides a simple REST API using the Gin framework to manage a list of movies.

## Code

```go
package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

// movie represents data about a film.
type movie struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Price    string `json:"price"`
}

// movies slice for demonstration
var movies = []movie{
	{ID: "1", Title: "The Dark Knight", Director: "Christopher Nolan", Price: "5.99"},
	{ID: "2", Title: "Tommy Boy", Director: "Peter Segal", Price: "2.99"},
	{ID: "3", Title: "The Shawshank Redemption", Director: "Frank Darabont", Price: "7.99"},
}

func main() {
	router := gin.New()
	router.LoadHTMLGlob("templates/*.html")
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewareFunc1, middlewareFunc2, middlewareFunc3())

	router.GET("/movie", getAllMovies)
	authRouter := router.Group("/auth", gin.BasicAuth(gin.Accounts{
		"Joe":   "baseball",
		"Kelly": "1234",
	}))

authRouter.GET("/movie", createMovieForm)
authRouter.POST("/movie", createMovie)

	router.Run(":8080")
}

func middlewareFunc1(c *gin.Context) {
	fmt.Println("middlewareFunc1 running")
	c.Next()
}

func middlewareFunc2(c *gin.Context) {
	fmt.Println("middlewareFunc2 running")
	c.Next()
}

func middlewareFunc3() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middlewareFunc3 running")
		c.Next()
	}
}

func getAllMovies(c *gin.Context) {
	c.HTML(http.StatusOK, "allmovies.html", movies)
}

func createMovieForm(c *gin.Context) {
	c.HTML(http.StatusOK, "createmovieform.html", nil)
}

func createMovie(c *gin.Context) {
	var newMovie movie
	newMovie.ID = c.PostForm("id")
	newMovie.Title = c.PostForm("title")
	newMovie.Director = c.PostForm("director")
	newMovie.Price = c.PostForm("price")
	movies = append(movies, newMovie)
	c.HTML(http.StatusOK, "allmovies.html", movies)
}
```

## API Endpoints

### 1. Get all Movies

**Request:**

```sh
curl -X GET http://localhost:8080/movie
```

**Response:**

```json
[
  {
    "id": "1",
    "title": "The Dark Knight",
    "director": "Christopher Nolan",
    "price": "5.99"
  },
  {
    "id": "2",
    "title": "Tommy Boy",
    "director": "Peter Segal",
    "price": "2.99"
  },
  {
    "id": "3",
    "title": "The Shawshank Redemption",
    "director": "Frank Darabont",
    "price": "7.99"
  }
]
```

### 2. Create a Movie Form (Requires Authentication)

**Request:**

```sh
curl -u Joe:baseball -X GET http://localhost:8080/auth/movie
```

### 3. Add a New Movie (Requires Authentication)

**Request:**

```sh
curl -u Joe:baseball -X POST http://localhost:8080/auth/movie \
  -d "id=4" -d "title=Inception" -d "director=Christopher Nolan" -d "price=6.99"
```

## Summary of API Calls

| HTTP Method | Endpoint      | Description                               |
| ----------- | ------------- | ----------------------------------------- |
| `GET`       | `/movie`      | Fetch all movies                          |
| `GET`       | `/auth/movie` | Load movie creation form                  |
| `POST`      | `/auth/movie` | Add a new movie (Authentication required) |

---

### Notes:

- `c.HTML` renders HTML templates from `templates/*.html`.
- Authentication is required for the `auth/movie` routes.
- The API runs on `localhost:8080`.

🚀 Enjoy coding with Gin!
