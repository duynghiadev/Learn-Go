# Gin Web Server API - cURL Commands

### **What does cURL stand for?**

**cURL** stands for **Client URL** and is a command-line tool used to transfer data to and from a server using various protocols (HTTP, HTTPS, FTP, etc.).

It is widely used for making API requests, downloading files, and testing web services.

## 1️⃣ GET - Hello World

```sh
curl -X GET http://localhost:8080/hello
```

## 2️⃣ GET - Greet Page

```sh
curl -X GET http://localhost:8080/greet
```

## 3️⃣ GET - Greet by Name

```sh
curl -X GET http://localhost:8080/greet/DuyNghia
```

## 4️⃣ GET - Many Data

```sh
curl -X GET http://localhost:8080/many
```

## 5️⃣ GET - Form Page

```sh
curl -X GET http://localhost:8080/form
```

## 6️⃣ POST - Submit Form

```sh
curl -X POST http://localhost:8080/form \
     -H "Content-Type: application/x-www-form-urlencoded" \
     -d "name=DuyNghia&food=Pizza"
```
