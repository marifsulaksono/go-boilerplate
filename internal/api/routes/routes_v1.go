package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/api/controller"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract"
)

type APIVersion1 struct {
	contract *contract.Contract
	e        *echo.Echo
	api      *echo.Group
}

func InitVersion(e *echo.Echo, path string, c *contract.Contract) APIVersion1 {
	return APIVersion1{
		c,
		e,
		e.Group(path),
	}
}

func (r *APIVersion1) UserAndAuth() {
	userController := controller.NewUserController(r.contract.Service.User)

	user := r.api.Group("/users")
	user.GET("", userController.Get)
}
