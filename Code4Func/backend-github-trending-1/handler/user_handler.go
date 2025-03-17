package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user":  "Ryan",
		"email": "ryan@gmail.com",
	})
}

func HandleSignUp(c echo.Context) error {
	type User struct {
		Email    string `json:"email"`
		FullName string `json:"name"`
		Age      int    `json:"age"`
	}

	user := User{
		Email:    "ryan@gmail.com",
		FullName: "Ryan Nguyen",
		Age:      23,
	}
	return c.JSON(http.StatusOK, user)
}
