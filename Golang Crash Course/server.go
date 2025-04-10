package main

import (
	"context"
	"fmt"
	"golang-rest-api/handler"
	"golang-rest-api/repository"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
)

func main() {
	// Initialize Firebase
	ctx := context.Background()
	sa := option.WithCredentialsFile("pragmatic-reviews.json") // Path to your .json file
	conf := &firebase.Config{
		// https://console.firebase.google.com/u/0/project/pragmatic-reviews-433b7/firestore/databases/-default-/data/~2Fposts~2F7fezUSvihRsbkI7B7364
		ProjectID: "pragmatic-reviews-433b7",
	}
	app, err := firebase.NewApp(ctx, conf, sa)

	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error initializing Firestore: %v\n", err)
	}
	defer client.Close()

	// Inject Firestore client into the repository
	postRepository := repository.NewPostRepository(client)

	// Initialize the router
	router := mux.NewRouter()
	const port string = ":8081"

	// Inject the repository into the handler
	postHandler := handler.NewPostHandler(postRepository)

	// Define the handler functions, passing in the repository
	router.HandleFunc("/", func(resp http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(resp, "Up and running...")
	})

	router.HandleFunc("/posts", postHandler.GetPosts).Methods("GET")
	router.HandleFunc("/posts", postHandler.AddPost).Methods("POST")

	log.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
