package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TodoItem struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func main() {
	godotenv.Load(".env.dev")

	dsn := os.Getenv("DB")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Connected to database:", db)

	now := time.Now().UTC()

	item := TodoItem{
		Id:          1,
		Title:       "Learn Go",
		Description: "Learn Go programming language",
		Status:      "Doing",
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": item,
		})
	})
	r.Run(":3000")
}
