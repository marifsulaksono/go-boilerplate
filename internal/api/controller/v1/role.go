package v1

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/api/controller/v1/dto"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract/service"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/helper"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/utils/response"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/service/interfaces"
)

type RoleController struct {
	Service interfaces.RoleService
}

func NewRoleController(svc service.ServiceContract) *RoleController {
	return &RoleController{
		Service: svc.GetRole(),
	}
}

func (h *RoleController) Get(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req dto.GetRoleRequest
	)

	data, err := h.Service.Get(ctx, req.ParseToModel())
	if err != nil {
		return response.BuildErrorResponse(c, err)
	}
	return response.BuildSuccessResponse(c, http.StatusOK, "Berhasil mendapatkan data role", data, nil)
}

func (h *RoleController) GetById(c echo.Context) error {
	var (
		ctx   = c.Request().Context()
		id, _ = uuid.Parse(c.Param("id"))
	)

	data, err := h.Service.GetById(ctx, id)
	if err != nil {
		return response.BuildErrorResponse(c, err)
	}
	return response.BuildSuccessResponse(c, http.StatusOK, "Berhasil mendapatkan data role", data, nil)
}

func (h *RoleController) Create(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		request dto.RoleRequest
	)

	if err := helper.BindRequest(c, &request, false); err != nil {
		return response.BuildErrorResponse(c, err)
	}

	data, err := h.Service.Create(ctx, request.ParseToModel())
	if err != nil {
		return response.BuildErrorResponse(c, err)
	}

	return response.BuildSuccessResponse(c, http.StatusCreated, "Berhasil menyimpan data role", map[string]string{"id": data}, nil)
}

func (h *RoleController) Update(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		id, _   = uuid.Parse(c.Param("id"))
		request dto.RoleRequest
	)

	if err := helper.BindRequest(c, &request, false); err != nil {
		return response.BuildErrorResponse(c, err)
	}

	data, err := h.Service.Update(ctx, request.ParseToModel(), id)
	if err != nil {
		return response.BuildErrorResponse(c, err)
	}

	return response.BuildSuccessResponse(c, http.StatusOK, "Berhasil memperbarui data role", map[string]string{"id": data}, nil)
}

func (h *RoleController) Delete(c echo.Context) error {
	var (
		ctx   = c.Request().Context()
		id, _ = uuid.Parse(c.Param("id"))
	)

	if err := h.Service.Delete(ctx, id); err != nil {
		return response.BuildErrorResponse(c, err)
	}

	return response.BuildSuccessResponse(c, http.StatusOK, "Berhasil menghapus data role", nil, nil)
}
