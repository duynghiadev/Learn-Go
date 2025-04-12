# Go REST API - Post Service (Clean Architecture)

## Description

This project implements a simple RESTful API service for managing blog posts. It is built using Go and leverages Firebase Firestore for data persistence. The primary goal of this project structure is to demonstrate the application of Clean Architecture principles for better maintainability, testability, and separation of concerns.

## Architecture Overview

The project follows the Clean Architecture pattern, dividing the application into distinct layers with specific dependency rules (dependencies only point inwards):

1. **Domain:** Contains the core business entities (`Post`) and logic. It has no dependencies on other layers.
2. **Usecase:** Orchestrates application-specific business logic. It defines interfaces for repositories and interacts with domain entities. Depends only on Domain.
3. **Interfaces:** Acts as adapters between the use cases and external systems. Includes HTTP handlers and repository implementations (e.g., Firestore adapter). Depends on Usecase and Domain.
4. **Infrastructure:** Contains external frameworks and tools like the database driver (Firebase SDK), web server (Gorilla Mux), and router setup. Depends on Interfaces.
5. **cmd:** Contains the main application entry point (`main.go`) responsible for initializing and wiring all the layers together (Dependency Injection).

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

1. **Clone the repository:**

   ```bash
   git clone <your-repo-url>
   cd golang-rest-api
   ```

2. **Ensure Go modules are downloaded:**

   ```bash
   go mod tidy
   ```

3. **Set Environment Variables (example for Linux/macOS):**

   ```bash
   export SERVER_PORT=":8081"
   export FIREBASE_PROJECT_ID="your-project-id"
   export FIREBASE_CREDENTIALS_PATH="pragmatic-reviews.json"
   ```

   _(Adjust `your-project-id` and the path to your credentials file)_

4. **Run the application:**

   ```bash
   go run ./cmd/api/main.go
   ```

5. The server should start, listening on the configured port (e.g., `http://localhost:8081`).

## API Endpoints

The API provides endpoints for managing posts. See the `API.md` document for detailed specifications.

- `GET /api/v1/posts`: Retrieve all posts.
- `POST /api/v1/posts`: Create a new post.

_(Add more endpoints here as they are implemented)_

# Dockerfile

Let me explain the key components of this Dockerfile:

1. Build Stage :

   - Uses `golang:1.21-alpine` as the base image for building
   - Sets up the working directory
   - Copies and downloads dependencies first (for better layer caching)
   - Copies the source code and builds the binary

2. Final Stage :

   - Uses `alpine:3.18` as the minimal base image
   - Copies only the necessary files (binary and credentials)
   - Exposes port 8081 (matching your server configuration)
   - Defines the command to run the application
     To build and run the container, use these commands:

To build and run the container, use these commands:

```bash
# Build the image
docker build -t golang-rest-api .

# Run the container
docker run -p 8081:8081 golang-rest-api
```

**Important Notes:**

1. Make sure your `pragmatic-reviews.json` file is in the project root directory
2. The container exposes port 8081 to match your application's configuration
3. The multi-stage build approach keeps the final image size small by excluding build tools and source code
   You might also want to add a `.dockerignore` file to exclude unnecessary files:

```.dockerignore
.git
.gitignore
README.md
*.md
.DS_Store
```

This setup provides a production-ready container for your Go REST API while maintaining a small image size and following best practices for containerization.
