# API Endpoints Documentation

# GitHub Trending API Documentation

Base URL: `http://localhost:3000`

## Authentication Endpoints

### Sign Up

- **Method**: `POST`
- **Endpoint**: `/user/sign-up`
- **Description**: Register a new user
- **Request Body**:
  ```json
  {
    "fullName": "string",
    "email": "string",
    "password": "string"
  }
  ```
- **Response**:
  ```json
  {
    "code": 200,
    "message": "Xử lý thành công",
    "data": {
      "fullName": "string",
      "email": "string",
      "role": "MEMBER"
    }
  }
  ```

### Sign In

- **Method**: `POST`
- **Endpoint**: `/user/sign-in`
- **Description**: Authenticate user and get access token
- **Admin Check**: Required (Only admin@gmail.com can access)
- **Request Body**:
  ```json
  {
    "email": "string",
    "password": "string"
  }
  ```
- **Response**:
  ```json
  {
    "code": 200,
    "message": "Xử lý thành công",
    "data": {
      "fullName": "string",
      "email": "string",
      "role": "MEMBER"
    }
  }
  ```
- **Error Response** (Non-admin):
  ```json
  {
    "code": 400,
    "message": "Bạn không không có quyền gọi api này !",
    "data": null
  }
  ```

## User Roles

The application supports two user roles:

1. **MEMBER**: Default role for regular users
2. **ADMIN**: Special role with additional privileges

## Database Schema

The application uses PostgreSQL with the following main tables:

### Users Table

```sql
CREATE TABLE "users" (
  "user_id" text PRIMARY KEY,
  "full_name" text,
  "email" text UNIQUE,
  "password" text,
  "role" text,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL
);
```

### Repos Table

```sql
CREATE TABLE "repos" (
  "name" text PRIMARY KEY,
  "description" text,
  "url" text,
  "color" text,
  "lang" text,
  "fork" text,
  "stars" text,
  "stars_today" text,
  "build_by" text,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL
);
```

### Bookmarks Table

```sql
CREATE TABLE "bookmarks" (
  "bid" text PRIMARY KEY,
  "user_id" text,
  "repo_name" text,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL
);
```

## Error Responses

The API may return the following error responses:

- `400 Bad Request`: Invalid request body, validation errors, or insufficient permissions
- `401 Unauthorized`: Invalid credentials
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: Resource not found
- `409 Conflict`: Resource already exists (e.g., email already registered)
- `500 Internal Server Error`: Server-side errors

## Authentication

The application uses JWT (JSON Web Tokens) for authentication. After successful sign-in, the token should be included in subsequent requests in the Authorization header:

```
Authorization: Bearer <token>
```

## Security Features

- Password hashing using bcrypt
- JWT-based authentication
- Input validation using validator package
- CORS enabled
- SQL injection prevention using prepared statements
- Admin role check middleware
