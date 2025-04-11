# 📘 API Documentation

**Base URL**: `http://localhost:8080/api/v1`

---

## 🔐 Authentication

### Register User

- **URL**: `/users/register`
- **Method**: `POST`
- **Content-Type**: `application/json`

#### Request Body:

```json
{
  "email": "test8@gmail.com",
  "firstName": "John",
  "lastName": "Doe",
  "password": "123456"
}
```

---

### Login

- **URL** : `/users/login`
- **Method** : `POST`
- **Content-Type** : `application/json`

#### Request Body:

```json
{
  "email": "test8@gmail.com",
  "password": "123456"
}
```

---

## 📁 Projects

> All project-related routes require an `Authorization` header with a valid Bearer token.

### Create Project

- **URL** : `/projects`
- **Method** : `POST`
- **Headers** :
  - `Authorization: Bearer <token>`
  - `Content-Type: application/json`

#### Request Body:

```json
{
  "name": "My New Project 3"
}
```

### Get Project by ID

- **URL** : `/projects/{id}`
- **Method** : `GET`
- **Headers** :
  - `Authorization: Bearer <token>`

### Get All Projects

- **URL** : `/projects`
- **Method** : `GET`
- **Headers** :
  - `Authorization: Bearer <token>`

### Delete Project

- **URL** : `/projects/{id}`
- **Method** : `DELETE`
- **Headers** :
  - `Authorization: Bearer <token>`

---

## ✅ Tasks

> All task-related routes require an `Authorization` header with a valid Bearer token.

### Create Task

- **URL** : `/tasks`
- **Method** : `POST`
- **Headers** :
  - `Authorization: Bearer <token>`
  - `Content-Type: application/json`

#### Request Body:

```json
{
  "name": "Implement Task API",
  "status": "TODO", // Options: TODO, IN_PROGRESS, IN_TESTING, DONE
  "projectID": 7,
  "assignedTo": 7
}
```

### Get Task by ID

- **URL** : `/tasks/{id}`
- **Method** : `GET`
- **Headers** :
  - `Authorization: Bearer <token>`

---

### Get All Tasks

- **URL** : `/tasks`
- **Method** : `GET`
- **Headers** :
  - `Authorization: Bearer <token>`

---

## 🧪 Notes

- Replace `<token>` with a valid JWT access token retrieved from the login endpoint.
- All protected endpoints require the `Authorization` header.
- You can test these APIs using Postman, Insomnia, or REST Client in VS Code (`.http` files).
