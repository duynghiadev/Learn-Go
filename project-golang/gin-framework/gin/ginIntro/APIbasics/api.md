# Movie API - cURL Commands

### **What does cURL stand for?**

**cURL** stands for **Client URL** and is a command-line tool used to transfer data to and from a server using various protocols (HTTP, HTTPS, FTP, etc.).

It is widely used for making API requests, downloading files, and testing web services.

## 1️⃣ Get All Movies

```sh
curl -X GET http://localhost:8080/movie
```

## 2️⃣ Get a Movie by ID

```sh
curl -X GET http://localhost:8080/movie/1
```

## 3️⃣ Create a New Movie

```sh
curl -X POST http://localhost:8080/movie \
     -H "Content-Type: application/json" \
     -d '{
          "id": "4",
          "title": "Inception",
          "director": "Christopher Nolan",
          "price": "6.99"
         }'
```

## 4️⃣ Update Movie Price

```sh
curl -X PATCH http://localhost:8080/movie/1
```

## 5️⃣ Delete a Movie

```sh
curl -X DELETE http://localhost:8080/movie/2
```
