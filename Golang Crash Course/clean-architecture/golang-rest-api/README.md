# Go REST API - Post Service (Clean Architecture)

## Description

This project implements a simple RESTful API service for managing blog posts. It is built using Go and leverages Firebase Firestore for data persistence. The primary goal of this project structure is to demonstrate the application of Clean Architecture principles for better maintainability, testability, and separation of concerns.

## Architecture Overview

The project follows the Clean Architecture pattern, dividing the application into distinct layers with specific dependency rules (dependencies only point inwards):

1.  **Domain:** Contains the core business entities (`Post`) and logic. It has no dependencies on other layers.
2.  **Usecase:** Orchestrates application-specific business logic. It defines interfaces for repositories and interacts with domain entities. Depends only on Domain.
3.  **Interfaces:** Acts as adapters between the use cases and external systems. Includes HTTP handlers and repository implementations (e.g., Firestore adapter). Depends on Usecase and Domain.
4.  **Infrastructure:** Contains external frameworks and tools like the database driver (Firebase SDK), web server (Gorilla Mux), and router setup. Depends on Interfaces.
5.  **cmd:** Contains the main application entry point (`main.go`) responsible for initializing and wiring all the layers together (Dependency Injection).

## Directory Structure

```structure
golang-rest-api/
├── cmd/
│ └── api/
│ └── main.go # Application entry point, wiring everything
├── domain/ # Core business logic and entities (innermost layer)
│ └── post.go
├── usecase/ # Application-specific business rules
│ ├── post_usecase.go # Defines the use case interface
│ └── post_interactor.go # Implements the use case interface
│ └── repository/ # Interfaces required by use cases (defined here!)
│ └── post_repository.go
├── interfaces/ # Adapters to interact with external systems
│ ├── handler/ # HTTP handlers (Controllers)
│ │ └── post_handler.go
│ └── repository/ # Repository implementations (Gateways)
│ └── firestore/
│ └── post_repository_firestore.go
├── infrastructure/ # Frameworks, drivers, external tools (outermost layer)
│ ├── datastore/
│ │ └── firebase.go # Firebase/Firestore initialization
│ └── router/
│ └── router.go # Router setup (e.g., Gorilla Mux)
├── go.mod
├── go.sum
└── pragmatic-reviews.json # Firebase credentials
```

## Prerequisites

- Go (version 1.18 or later recommended)
- A Google Cloud Platform project with Firestore enabled.
- Firebase Admin SDK credentials file (`.json`).

## Configuration

The application requires the following configuration, typically loaded via environment variables (see `cmd/api/main.go` `loadConfig` function):

- `SERVER_PORT`: The port the HTTP server listens on (e.g., `:8081`).
- `FIREBASE_PROJECT_ID`: Your Google Cloud/Firebase Project ID.
- `FIREBASE_CREDENTIALS_PATH`: The file path to your Firebase Admin SDK credentials JSON file.

Place your Firebase credentials file (e.g., `pragmatic-reviews.json`) in the project root or specify its path via the environment variable.

## Running the Application

1.  **Clone the repository:**
    ```bash
    git clone <your-repo-url>
    cd golang-rest-api
    ```
2.  **Ensure Go modules are downloaded:**
    ```bash
    go mod tidy
    ```
3.  **Set Environment Variables (example for Linux/macOS):**
    ```bash
    export SERVER_PORT=":8081"
    export FIREBASE_PROJECT_ID="your-project-id"
    export FIREBASE_CREDENTIALS_PATH="pragmatic-reviews.json"
    ```
    _(Adjust `your-project-id` and the path to your credentials file)_
4.  **Run the application:**
    ```bash
    go run ./cmd/api/main.go
    ```
5.  The server should start, listening on the configured port (e.g., `http://localhost:8081`).

## API Endpoints

The API provides endpoints for managing posts. See the `API.md` document for detailed specifications.

- `GET /api/v1/posts`: Retrieve all posts.
- `POST /api/v1/posts`: Create a new post.

_(Add more endpoints here as they are implemented)_
