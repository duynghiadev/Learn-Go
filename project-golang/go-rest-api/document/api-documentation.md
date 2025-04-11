# 📘 REST API Documentation

## 🔗 Base URL

```
http://localhost:8080
```

---

## 🔐 Authentication

- **Bearer Token** authentication is required for protected endpoints.
- Token can be obtained from the **Login** endpoint.
- Add token in:
  - **Authorization Header**: `Bearer <token>`
  - **Query Parameter**: `?token=<token>`

---

## 📌 Endpoints

### 🧍 Authentication

subrouter: `/api/v1`

#### ✅ Register User

**POST** `/users/register`
**Request Body (JSON):**

```json
{
  "email": "user@example.com",
  "firstName": "John",
  "lastName": "Doe",
  "password": "yourpassword"
}
```

**Response:**

```json
{
  "id": 1,
  "email": "user@example.com",
  "firstName": "John",
  "lastName": "Doe",
  "createdAt": "2024-01-20T10:00:00Z"
}
```

---

#### 🔑 Login

**POST** `/login`
**Request Body (JSON):**

```json
{
  "email": "user@example.com",
  "password": "yourpassword"
}
```

**Response:**

```json
{
  "token": "eyJhbGciOiJIUzI1NiIs..."
}
```

---

### 📂 Projects

#### ➕ Create Project

**POST** `/projects`
**Headers:**

```
Authorization: Bearer <token>
```

**Request Body (JSON):**

```json
{
  "name": "My New Project"
}
```

**Response:**

```json
{
  "id": 1,
  "name": "My New Project",
  "createdAt": "2024-01-20T10:00:00Z"
}
```

---

#### 📄 Get All Projects

**GET** `/projects`
**Headers:**

```
Authorization: Bearer <token>
```

**Response:**

```json
[
  {
    "id": 1,
    "name": "My New Project",
    "createdAt": "2024-01-20T10:00:00Z"
  }
]
```

---

#### 📄 Get Project by ID

**GET** `/projects/{id}`
**Headers:**

```
Authorization: Bearer <token>
```

**Response:**

```json
{
  "id": 1,
  "name": "My New Project",
  "createdAt": "2024-01-20T10:00:00Z"
}
```

---

### ✅ Tasks

#### ➕ Create Task

**POST** `/projects/{projectId}/tasks`
**Headers:**

```
Authorization: Bearer <token>
```

**Request Body (JSON):**

```json
{
  "name": "New Task",
  "assignedToId": 1,
  "status": "TODO"
}
```

**Response:**

```json
{
  "id": 1,
  "name": "New Task",
  "status": "TODO",
  "projectId": 1,
  "assignedToId": 1,
  "createdAt": "2024-01-20T10:00:00Z"
}
```

---

#### 📄 Get Project Tasks

**GET** `/projects/{projectId}/tasks`
**Headers:**

```
Authorization: Bearer <token>
```

**Response:**

```json
[
  {
    "id": 1,
    "name": "New Task",
    "status": "TODO",
    "projectId": 1,
    "assignedToId": 1,
    "createdAt": "2024-01-20T10:00:00Z"
  }
]
```

---

#### 🔄 Update Task Status

**PUT** `/projects/{projectId}/tasks/{taskId}`
**Headers:**

```
Authorization: Bearer <token>
```

**Request Body (JSON):**

```json
{
  "status": "IN_PROGRESS"
}
```

**Response:**

```json
{
  "id": 1,
  "name": "New Task",
  "status": "IN_PROGRESS",
  "projectId": 1,
  "assignedToId": 1,
  "createdAt": "2024-01-20T10:00:00Z"
}
```

---

## ⚠️ Status Codes

- `200` OK: Success
- `201` Created
- `400` Bad Request
- `401` Unauthorized
- `404` Not Found
- `500` Internal Server Error

---

## 🏷️ Task Status Values

- `TODO`
- `IN_PROGRESS`
- `IN_TESTING`
- `DONE`

---

## 🧪 Postman Collection Setup

### 📁 Create a Collection in Postman

#### 🔧 Set Environment Variables:

| Key        | Value                        |
| ---------- | ---------------------------- |
| `BASE_URL` | `http://localhost:8080`      |
| `TOKEN`    | Your JWT token from `/login` |

---

### 📂 Create the Following Requests

#### **Authentication Folder**

- **Register**: `POST {{BASE_URL}}/register`
- **Login**: `POST {{BASE_URL}}/login`

#### **Projects Folder**

- **Create Project**: `POST {{BASE_URL}}/projects`
- **Get All Projects**: `GET {{BASE_URL}}/projects`
- **Get Project by ID**: `GET {{BASE_URL}}/projects/{{project_id}}`

#### **Tasks Folder**

- **Create Task**: `POST {{BASE_URL}}/projects/{{project_id}}/tasks`
- **Get Tasks**: `GET {{BASE_URL}}/projects/{{project_id}}/tasks`
- **Update Task**: `PUT {{BASE_URL}}/projects/{{project_id}}/tasks/{{task_id}}`

---

### 🔐 Setup Authorization in Postman

- Go to **Collection Settings → Authorization**
- Choose **Bearer Token**
- Set Token: `{{TOKEN}}`

---

### ✅ Test Flow

1. Register a new user
2. Login and save the token
3. Create a project
4. Create tasks
5. Update task status
