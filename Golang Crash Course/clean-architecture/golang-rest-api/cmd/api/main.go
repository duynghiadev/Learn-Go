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
)

func main() {
	// --- Configuration ---
	projectID := "pragmatic-reviews-433b7"
	credentialsPath := "pragmatic-reviews.json"
	serverPort := ":8081"

	// Add Redis configuration
	redisHost := "localhost"
	redisPort := "6379"
	redisPassword := "" // Set your password if needed

	// --- Initialization Context ---
	// Use a background context for setup, potentially with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Initialize Redis
	log.Println("Main: Initializing Redis...")
	redisClient, err := datastore.InitRedis(ctx, redisHost, redisPort, redisPassword)
	if err != nil {
		log.Fatalf("Main: Failed to initialize Redis: %v", err)
		os.Exit(1)
	}
	defer redisClient.Close()

	// --- Infrastructure Initialization ---
	log.Println("Main: Initializing Firestore...")
	firestoreClient, err := datastore.InitFirestore(ctx, projectID, credentialsPath)
	if err != nil {
		log.Fatalf("Main: Failed to initialize Firestore: %v", err)
		os.Exit(1)
	}
	// defer firestoreClient.Close() // Can cause issues if context used for init is cancelled early

	// --- Dependency Injection (Wiring the Layers) ---
	log.Println("Main: Wiring application layers...")

	// 1. Repository Layer (Concrete Implementation)
	// We need the Firestore implementation which satisfies the usecase.repository.PostRepository interface.
	postRepo := firestoreRepo.NewFirestorePostRepository(firestoreClient)

	// 2. Use Case Layer (Concrete Implementation - Interactor)
	// The interactor needs a PostRepository interface, we provide the Firestore implementation.
	postUseCase := usecase.NewPostInteractor(postRepo)

	// 3. Interface Adapter Layer (HTTP Handler)
	// The handler needs a PostUseCase interface, we provide the interactor implementation.
	postHandler := handler.NewPostHandler(postUseCase)

	// 4. Router (Infrastructure)
	// The router needs the concrete handlers to map routes.
	router := infraRouter.NewRouter(postHandler)

	// --- Start Server ---
	log.Printf("Main: Starting HTTP server on port %s", serverPort)
	server := &http.Server{
		Addr:    serverPort,
		Handler: router,
		// Add timeouts for production readiness
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Main: Could not listen on %s: %v\n", serverPort, err)
	}

	log.Println("Main: Server stopped gracefully.")
}
