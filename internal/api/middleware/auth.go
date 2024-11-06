package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/helper"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/utils/response"
)

/*
	this middleware is for verifying token
	you can use this middleware in your route or group/global

	how to use it
	1. import middleware
	2. use middleware.JWTMiddleware

	more info contact me @marifsulaksono
*/

func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == "" {
				err := response.NewCustomError(http.StatusUnauthorized, "Invalid token", nil)
				return response.BuildErrorResponse(c, err)
			}
			user, err := helper.VerifyTokenJWT(tokenString, false)
			if err != nil {
				err := response.NewCustomError(http.StatusUnauthorized, "Invalid token", nil)
				return response.BuildErrorResponse(c, err)
			}

			c.Set("user_id", user.ID) // set saves data in the context

			return next(c)
		}
	}
}
