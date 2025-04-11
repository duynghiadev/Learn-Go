# Golang Transaction Project

A demonstration of multi-layer architecture using Golang and GORM for database transactions.

## 📂 Project Structure

This project implements a clean architecture approach with the following components:

- **Model**: Handles database connections and entity definitions
- **Route**: Contains API route configurations
- **Controller**: Handles business logic
- **Repository**: Manages data access operations

## 🚀 Setup and Configuration

### Prerequisites

- Go installed on your system
- MySQL/PostgreSQL database server

### Database Connection

The project uses GORM for database operations. Database connection is established in the `model` package.

### Main Application

The entry point of the application (`main.go`):

```go
package main

import (
    "golang-transaction/model"
    "golang-transaction/route"
)

func main() {
    db, _ := model.DBConnection()
    route.SetupRoutes(db)
}
```

## 🔧 Running the Application

1. Clone the repository
2. Configure your database settings
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Run the application:
   ```bash
   go run main.go
   ```

## 🌐 API Endpoints

The application exposes the following RESTful endpoints:

### User Management

#### Create User

- **URL**: `http://localhost:8080/users`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "email": "test2@gmail.com",
    "wallet": 2000
  }
  ```

#### Get All Users

- **URL**: `http://localhost:8080/users`
- **Method**: `GET`

### Money Transfer

#### Create Money Transfer

- **URL**: `http://localhost:8080/money-transfer`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "receiver": 5,
    "giver": 1,
    "amount": 100.5
  }
  ```

## 📦 Project Dependencies

- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [GORM](https://gorm.io/) - ORM library for database operations

## 🔍 Architecture Overview

The application follows a layered architecture pattern:

```
├── model/            # Database models and connections
├── repository/       # Data access layer
├── controller/       # Business logic
├── route/            # API routes
└── main.go           # Application entry point
```

## 🔐 Transaction Management

The application demonstrates proper transaction handling with GORM to ensure data integrity during money transfers between user wallets.
