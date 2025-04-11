# User Authentication API (Gin + MySQL)

## Overview

This API provides user authentication (login/logout) using Gin, MySQL, and session-based authentication.

## Endpoints

### 1️⃣ Home Page

#### GET `/`

Displays the homepage.

**Request:**

```sh
curl -X GET http://localhost:8080/
```

**Response:**

```html
<html>
  Home Page
</html>
```

---

### 2️⃣ Login Page

#### GET `/login`

Displays the login form.

**Request:**

```sh
curl -X GET http://localhost:8080/login
```

**Response:**

```html
<html>
  Login Form
</html>
```

---

### 3️⃣ Login (Authenticate User)

#### POST `/login`

Verifies username and password. Starts a session if login is successful.

**Request:**

```sh
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "username=testuser&password=password123"
```

**Success Response (200 OK):**

```html
<html>
  Welcome testuser
</html>
```

**Error Response (401 Unauthorized):**

```html
<html>
  Check username and password
</html>
```

---

### 4️⃣ Get User Profile (Requires Authentication)

#### GET `/user/profile`

Retrieves logged-in user profile. Requires session authentication.

**Request:**

```sh
curl -X GET http://localhost:8080/user/profile \
  --cookie "session=<your_session_cookie>"
```

**Success Response (200 OK):**

```html
<html>
  Profile of testuser
</html>
```

**Error Response (403 Forbidden - Not Logged In):**

```html
<html>
  Please login first
</html>
```

---

## Session Handling

- Sessions are stored as cookies.
- Middleware `auth` ensures only logged-in users can access `/user/profile`.

---

## Database Schema

```sql
CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    pswd_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    active BOOLEAN DEFAULT TRUE,
    ver_hash VARCHAR(255),
    timeout TIMESTAMP
);
```
