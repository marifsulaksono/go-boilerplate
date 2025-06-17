package v1

import (
	"github.com/labstack/echo/v4"
	controller "github.com/marifsulaksono/go-echo-boilerplate/internal/api/controller/v1"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/api/middleware"
)

func AuthRoutes(api *echo.Group, ctrl *controller.AuthController) {
	auth := api.Group("/auth")

	auth.POST("/register", ctrl.Register)
	auth.POST("/login", ctrl.Login, middleware.RateLimitMiddleware(5, 300)) // limit to 5 requests per 5 minutes
	auth.POST("/new-access-token", ctrl.RefreshAccessToken)
	auth.POST("/logout", ctrl.Logout)
}
