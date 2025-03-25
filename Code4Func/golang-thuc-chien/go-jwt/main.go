package main

import (
	"fmt"
	"net/http"

	"github.com/duynghiadev/go-jwt/config"
	"github.com/duynghiadev/go-jwt/driver"
	"github.com/duynghiadev/go-jwt/handler"
)

func main() {
	driver.ConnectMongoDB(config.DB_USER, config.DB_PASS)

	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/user", handler.GetUser)

	fmt.Println("Server running [:9999]")
	http.ListenAndServe(":9999", nil)
}
