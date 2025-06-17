package v1

import (
	"github.com/labstack/echo/v4"
	controller "github.com/marifsulaksono/go-echo-boilerplate/internal/api/controller/v1"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/api/middleware"
)

func UserRoutes(api *echo.Group, ctrl *controller.UserController) {
	user := api.Group("/users")
	user.Use(middleware.JWTMiddleware()) // use middleware jwt general on user routes

	user.GET("", ctrl.Get)
	user.GET("/:id", ctrl.GetById)
	user.POST("", ctrl.Create)
	user.PUT("/:id", ctrl.Update)
	user.DELETE("/:id", ctrl.Delete)
}
