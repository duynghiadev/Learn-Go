package main

import (
	"github.com/duynghiadev/backend-github-trending/db"
	"github.com/duynghiadev/backend-github-trending/handler"
	"github.com/labstack/echo"
)

func main() {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "postgres",
		Password: "duynghia123",
		DbName:   "code4func",
	}

	sql.Connect()
	defer sql.Close()

	e := echo.New()
	e.GET("/", handler.Welcome)

	e.GET("/user/sign-in", handler.HandleSignIn)
	e.GET("/user/sign-up", handler.HandleSignUp)

	e.Logger.Fatal(e.Start(":9999"))
}
