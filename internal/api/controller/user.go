package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/service"
)

type UserController struct {
	Service service.UserService
}

func NewUserController(s service.UserService) *UserController {
	return &UserController{
		Service: s,
	}
}

func (h *UserController) Get(c echo.Context) error {
	var (
		ctx = c.Request().Context()
	)
	str := h.Service.Get(ctx)
	return c.JSON(200, str)
}
