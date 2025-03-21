package main

import (
	"backend-github-trending/db"
	_ "backend-github-trending/docs"
	"backend-github-trending/handler"
	"backend-github-trending/helper"
	"backend-github-trending/log"
	"backend-github-trending/repository/repo_impl"
	"backend-github-trending/router"
	"fmt"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func init() {
	fmt.Println("DEV ENVIRONMENT")
	os.Setenv("APP_NAME", "github")
	log.InitLogger(false)
}

// @title Github Trending API
// @version 1.0
// @description More
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey jwt
// @in header
// @name Authorization

// @host localhost:3002
// @BasePath /
func main() {
	sql := &db.Sql{
		Host:     "localhost", //localhost,host.docker.internal
		Port:     5432,
		UserName: "postgres",
		Password: "duynghia123",
		DbName:   "code4func",
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	structValidator := helper.NewStructValidator()
	structValidator.RegisterValidate()

	e.Validator = structValidator

	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}

	repoHandler := handler.RepoHandler{
		GithubRepo: repo_impl.NewGithubRepo(sql),
	}

	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
		RepoHandler: repoHandler,
	}
	api.SetupRouter()

	go scheduleUpdateTrending(360*time.Second, repoHandler)

	e.Logger.Fatal(e.Start(":3002"))
}

func scheduleUpdateTrending(timeSchedule time.Duration, handler handler.RepoHandler) {
	ticker := time.NewTicker(timeSchedule)
	go func() {
		for range ticker.C {
			fmt.Println("Checking from github...")
			helper.CrawlRepo(handler.GithubRepo)
		}
	}()
}
