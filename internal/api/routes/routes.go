package routes

import (
	"github.com/labstack/echo/v4"
	controllerv1 "github.com/marifsulaksono/go-echo-boilerplate/internal/api/controller/v1"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/api/middleware"
	v1 "github.com/marifsulaksono/go-echo-boilerplate/internal/api/routes/v1"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract"
)

type APIVersion struct {
	Contract *contract.Contract
	Echo     *echo.Echo
	Api      *echo.Group
}

func InitVersion(e *echo.Echo, path string, c *contract.Contract) APIVersion {
	return APIVersion{
		Contract: c,
		Echo:     e,
		Api:      e.Group(path),
	}
}

func RouteV1(av *APIVersion) {
	userController := controllerv1.NewUserController(av.Contract.Service)
	authController := controllerv1.NewAuthController(av.Contract.Service)
	roleController := controllerv1.NewRoleController(av.Contract.Service)

	av.Api.Use(middleware.LogMiddleware) // use middleware logger
	v1.UserRoutes(av.Api, userController)
	v1.AuthRoutes(av.Api, authController)
	v1.RoleRoutes(av.Api, roleController)
}
