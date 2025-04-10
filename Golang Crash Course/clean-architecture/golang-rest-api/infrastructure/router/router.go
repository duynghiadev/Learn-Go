package router

import (
	"fmt"
	"golang-clean-architecture/interfaces/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(postHandler *handler.PostHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "API is up and running...")
	}).Methods("GET")

	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/posts", postHandler.GetPosts).Methods("GET")
	api.HandleFunc("/posts", postHandler.AddPost).Methods("POST")
	// Add other post routes here (e.g., GET /posts/{id}, PUT /posts/{id}, DELETE /posts/{id})

	log.Println("Router: Routes configured")
	return router
}
