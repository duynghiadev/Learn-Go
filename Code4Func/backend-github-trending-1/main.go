package main

import (
	"fmt"
	"os"
	"time"

	"github.com/duynghiadev/backend-github-trending/db"
	"github.com/duynghiadev/backend-github-trending/handler"
	"github.com/duynghiadev/backend-github-trending/helper"
	"github.com/duynghiadev/backend-github-trending/log"
	"github.com/duynghiadev/backend-github-trending/repository/repo_impl"
	"github.com/duynghiadev/backend-github-trending/router"
	"github.com/labstack/echo"
)

func init() {
	os.Setenv("APP_NAME", "github")
	log.InitLogger(false)
}

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

	go scheduleUpdateTrending(15*time.Second, repoHandler)

	e.Logger.Fatal(e.Start(":3001"))
}

func scheduleUpdateTrending(timeSchedule time.Duration, handler handler.RepoHandler) {
	ticker := time.NewTicker(timeSchedule)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Checking from github...")
				helper.CrawlRepo(handler.GithubRepo)
			}
		}
	}()
}
