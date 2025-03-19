package middleware

import (
	"github.com/duynghiadev/backend-github-trending/model"
	"github.com/duynghiadev/backend-github-trending/security"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte(security.JWT_KEY),
	}

	return middleware.JWTWithConfig(config)
}
