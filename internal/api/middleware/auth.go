package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/helper"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == "" {
				return c.JSON(401, "Invalid token")
			}
			user, err := helper.VerifyTokenJWT(tokenString, false)
			if err != nil {
				return c.JSON(401, "Invalid token")
			}

			c.Set("user", user) // set saves data in the context

			return next(c)
		}
	}
}
