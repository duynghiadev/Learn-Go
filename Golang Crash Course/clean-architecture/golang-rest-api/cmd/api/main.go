package main

import (
	"context"
	"golang-clean-architecture/infrastructure/datastore"
	infraRouter "golang-clean-architecture/infrastructure/router"
	"golang-clean-architecture/interfaces/handler"
	firestoreRepo "golang-clean-architecture/interfaces/repository/firestore"
	"golang-clean-architecture/usecase"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	projectID := "pragmatic-reviews-433b7"
	credentialsPath := "pragmatic-reviews.json"
	serverPort := ":8081"

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	log.Println("Main: Initializing Redis...")
	redisClient, err := datastore.InitRedis(ctx, redisHost, redisPort, redisPassword)
	if err != nil {
		log.Fatalf("Main: Failed to initialize Redis: %v", err)
		os.Exit(1)
	}
	defer redisClient.Close()

	log.Println("Main: Initializing Firestore...")
	firestoreClient, err := datastore.InitFirestore(ctx, projectID, credentialsPath)
	if err != nil {
		log.Fatalf("Main: Failed to initialize Firestore: %v", err)
		os.Exit(1)
	}
	log.Println("Main: Wiring application layers...")

	postRepo := firestoreRepo.NewFirestorePostRepository(firestoreClient, redisClient)
	postUseCase := usecase.NewPostInteractor(postRepo)
	postHandler := handler.NewPostHandler(postUseCase)

	router := infraRouter.NewRouter(postHandler)

	log.Printf("Main: Starting HTTP server on port %s", serverPort)
	server := &http.Server{
		Addr:    serverPort,
		Handler: router,

		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Main: Could not listen on %s: %v\n", serverPort, err)
	}

	log.Println("Main: Server stopped gracefully.")
}
