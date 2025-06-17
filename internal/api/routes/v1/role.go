package v1

import (
	"github.com/labstack/echo/v4"
	controller "github.com/marifsulaksono/go-echo-boilerplate/internal/api/controller/v1"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/api/middleware"
)

func RoleRoutes(api *echo.Group, ctrl *controller.RoleController) {
	role := api.Group("/roles")
	role.Use(middleware.JWTMiddleware()) // use middleware jwt general on role routes

	role.GET("", ctrl.Get)
	role.GET("/:id", ctrl.GetById)
	role.POST("", ctrl.Create)
	role.PUT("/:id", ctrl.Update)
	role.DELETE("/:id", ctrl.Delete)
}
