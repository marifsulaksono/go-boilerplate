package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/api/dto"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/helper"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/service"
)

type AuthController struct {
	Service service.AuthService
}

func NewAuthController(s service.AuthService) *AuthController {
	return &AuthController{
		Service: s,
	}
}

func (h *AuthController) Login(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		request dto.LoginRequest
	)

	if err := helper.BindRequest(c, &request, false); err != nil {
		return err
	}

	data, err := h.Service.Login(ctx, request.ParseToModel())
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, data)
}
