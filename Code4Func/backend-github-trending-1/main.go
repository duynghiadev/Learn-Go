package main

import (
	"fmt"
	"os"

	"github.com/duynghiadev/backend-github-trending/db"
	"github.com/duynghiadev/backend-github-trending/handler"
	"github.com/duynghiadev/backend-github-trending/log"
	"github.com/duynghiadev/backend-github-trending/repository/repo_impl"
	"github.com/duynghiadev/backend-github-trending/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	fmt.Println("init package main")
	os.Setenv("APP_NAME", "github")
	log.InitLogger(false)
}

func main() {
	fmt.Println("main function")
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

	e.Use(middleware.AddTrailingSlash())

	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}

	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
	}
	api.SetupRouter()

	e.Logger.Fatal(e.Start(":3001"))
}
