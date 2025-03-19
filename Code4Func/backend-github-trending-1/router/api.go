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
	api.Echo.POST("/user/sign-in", api.UserHandler.HandleSignIn, middleware.ISAdmin())

	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)
}
