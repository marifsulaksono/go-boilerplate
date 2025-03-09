package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/api/dto"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/helper"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/utils/response"
	service "github.com/marifsulaksono/go-echo-boilerplate/internal/service/interfaces"
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
		ip      = c.RealIP()
		request dto.LoginRequest
	)

	if err := helper.BindRequest(c, &request, false); err != nil {
		return response.BuildErrorResponse(c, err)
	}

	data, err := h.Service.Login(ctx, request.ParseToModel(), ip)
	if err != nil {
		return response.BuildErrorResponse(c, err)
	}

	return response.BuildSuccessResponse(c, http.StatusOK, "Login berhasil", data)
}

func (h *AuthController) RefreshAccessToken(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		request dto.RefreshAccessTokenRequest
	)

	if err := helper.BindRequest(c, &request, false); err != nil {
		return response.BuildErrorResponse(c, err)
	}

	data, err := h.Service.RefreshAccessToken(ctx, request.RefreshToken)
	if err != nil {
		return response.BuildErrorResponse(c, err)
	}

	return response.BuildSuccessResponse(c, http.StatusOK, "Berhasil mendapatkan access token baru", data)
}

func (h *AuthController) Logout(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		request dto.RefreshAccessTokenRequest
	)

	if err := helper.BindRequest(c, &request, false); err != nil {
		return response.BuildErrorResponse(c, err)
	}

	err := h.Service.Logout(ctx, request.RefreshToken)
	if err != nil {
		return response.BuildErrorResponse(c, err)
	}

	return response.BuildSuccessResponse(c, http.StatusOK, "Logout berhasil", nil)
}
