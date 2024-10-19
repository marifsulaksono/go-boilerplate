package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract"
)

type APIVersion struct {
	contract *contract.Contract
	e        *echo.Echo
	api      *echo.Group
}

func InitVersion(e *echo.Echo, path string, c *contract.Contract) APIVersion {
	return APIVersion{
		c,
		e,
		e.Group(path),
	}
}
