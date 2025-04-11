# Go JWT API Documentation

This document provides information about all available API endpoints in the Go JWT project.

## Base URL

```
http://localhost:9999
```

## Authentication

Most endpoints require JWT authentication. Include the JWT token in the Authorization header:

```
Authorization: Bearer <your_jwt_token>
```

## Endpoints

### 1. Register User

Register a new user in the system.

**Endpoint:** `POST /register`

**Request Body:**

```json
{
  "email": "string",
  "password": "string",
  "displayName": "string"
}
```

**Response:**

```json
{
  "token": "string",
  "status": 200
}
```

**Status Codes:**

- 200: Success
- 400: Bad Request
- 409: Conflict (Email already exists)
- 500: Internal Server Error

### 2. Login User

Authenticate a user and get a JWT token.

**Endpoint:** `POST /login`

**Request Body:**

```json
{
  "email": "string",
  "password": "string"
}
```

**Response:**

```json
{
  "token": "string",
  "status": 200
}
```

**Status Codes:**

- 200: Success
- 400: Bad Request
- 401: Unauthorized
- 500: Internal Server Error

### 3. Get User Information

Get the current user's information using their JWT token.

**Endpoint:** `GET /user`

**Headers:**

```
Authorization: Bearer <your_jwt_token>
```

**Response:**

```json
{
    "email": "string",
    "displayName": "string",
    "exp": number
}
```

**Status Codes:**

- 200: Success
- 403: Forbidden (No token provided)
- 500: Internal Server Error

## Notes

- All requests and responses are in JSON format
- The JWT token expires after 120 seconds
- The server runs on port 9999 by default
