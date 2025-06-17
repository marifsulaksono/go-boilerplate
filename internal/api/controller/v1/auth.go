package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/api/controller/v1/dto"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract/service"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/helper"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/utils/response"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/service/interfaces"
)

type AuthController struct {
	Service interfaces.AuthService
}

func NewAuthController(svc service.ServiceContract) *AuthController {
	return &AuthController{
		Service: svc.GetAuth(),
	}
}

func (h *AuthController) Register(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		request dto.RegisterRequest
	)

	if err := helper.BindRequest(c, &request, false); err != nil {
		return response.BuildErrorResponse(c, err)
	}

	err := h.Service.Register(ctx, request.ParseToModel())
	if err != nil {
		return response.BuildErrorResponse(c, err)
	}

	return response.BuildSuccessResponse(c, http.StatusOK, "Register berhasil", nil, nil)
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

	data, err := h.Service.Login(ctx, request.ParseToModel(ip))
	if err != nil {
		return response.BuildErrorResponse(c, err)
	}

	return response.BuildSuccessResponse(c, http.StatusOK, "Login berhasil", data, nil)
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

	return response.BuildSuccessResponse(c, http.StatusOK, "Berhasil mendapatkan access token baru", data, nil)
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

	return response.BuildSuccessResponse(c, http.StatusOK, "Logout berhasil", nil, nil)
}
