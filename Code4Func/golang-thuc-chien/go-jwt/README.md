# Go JWT Authentication Project

A practical implementation of JWT (JSON Web Token) authentication in Go, following the repository pattern.

## 📚 Learning Resources

This project is part of the [Golang Thực Chiến](https://www.youtube.com/playlist?list=PLVDJsRQrTUz7-fSzZWtF726st5AYQc57A) tutorial series.

## 🏗️ Project Structure

This project implements the repository pattern for better code organization and maintainability.

### 📁 Project Tree

```
go-jwt/
├── API_DOCUMENTATION.md
├── README.md
├── config/           # Configuration files
├── driver/          # Database driver setup
├── handler/         # HTTP request handlers
├── main.go          # Application entry point
├── model/           # Data models and DTOs
│   ├── error_user.go
│   ├── req_login.go
│   ├── req_reg.go
│   ├── res_err.go
│   ├── res_reg.go
│   └── user.go
├── repository/      # Repository layer
│   ├── repoimpl/   # Repository implementations
│   └── user_repo.go
├── go.mod
└── go.sum
```

## 🏛️ Architecture

This project follows a clean architecture approach with the following key components:

### Repository Pattern

- **Purpose**: Abstracts data persistence operations
- **Implementation**:
  - `repository/user_repo.go`: Defines the interface for user operations
  - `repository/repoimpl/`: Contains concrete implementations of the repository interfaces
  - Separates data access logic from business logic

### Layered Architecture

1. **Handler Layer** (`handler/`)

   - Handles HTTP requests and responses
   - Input validation and request processing
   - Routes requests to appropriate business logic

2. **Repository Layer** (`repository/`)

   - Abstracts database operations
   - Provides data access interfaces
   - Implements data persistence logic

3. **Model Layer** (`model/`)

   - Defines data structures and DTOs
   - Contains business entities
   - Includes request/response models

4. **Driver Layer** (`driver/`)

   - Manages database connections
   - Handles external service integrations
   - Provides infrastructure setup

5. **Config Layer** (`config/`)
   - Manages application configuration
   - Handles environment variables
   - Provides configuration interfaces

### Key Design Principles

- **Separation of Concerns**: Each layer has a specific responsibility
- **Dependency Inversion**: High-level modules don't depend on low-level modules
- **Interface-based Design**: Uses interfaces for loose coupling
- **Clean Architecture**: Follows SOLID principles for maintainable code

## 🔧 Database Configuration

The project uses MongoDB as its database. Here's the connection string:

```
mongodb+srv://duynghia22302:duynghia123@go-jwt.nt5ml.mongodb.net/?retryWrites=true&w=majority&appName=go-jwt
```

## ⚠️ Security Note

Please note that the database credentials are exposed in this README. In a production environment, you should:

1. Use environment variables
2. Never commit sensitive credentials to version control
3. Use a secure secrets management system

## 🚀 Getting Started

(Add installation and setup instructions here)

## 📝 License

(Add license information here)
