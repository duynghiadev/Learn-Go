package router

import (
	"github.com/duynghiadev/backend-github-trending/handler"
	"github.com/duynghiadev/backend-github-trending/middleware"
	"github.com/labstack/echo"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetupRouter() {
	api.Echo.POST("/user/sign-in", api.UserHandler.HandleSignIn)

	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)

	user := api.Echo.Group("/user", middleware.JWTMiddleware())
	user.GET("/profile", api.UserHandler.Profile)
	user.PUT("/profile/update", api.UserHandler.UpdateProfile)
}
