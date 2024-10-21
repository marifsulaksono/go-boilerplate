package controller

import (
	"errors"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/api/dto"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/service"
	"gorm.io/gorm"
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

	data, err := h.Service.Get(ctx)
	if err != nil {
		return c.JSON(500, err.Error())
	}
	return c.JSON(200, data)
}

func (h *UserController) GetById(c echo.Context) error {
	var (
		ctx   = c.Request().Context()
		id, _ = uuid.Parse(c.Param("id"))
	)

	data, err := h.Service.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(404, "Data Not Found")
		}
		return c.JSON(500, err.Error())
	}
	return c.JSON(200, data)
}

func (h *UserController) Create(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		request dto.UserRequest
	)

	if err := c.Bind(&request); err != nil {
		return c.JSON(400, err.Error())
	}

	if err := c.Validate(&request); err != nil {
		return c.JSON(400, err.Error())
	}

	user, err := h.Service.Create(ctx, request.ParseToModel())
	if err != nil {
		return c.JSON(400, err.Error())
	}

	return c.JSON(201, user)
}

func (h *UserController) Update(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		id, _   = uuid.Parse(c.Param("id"))
		request dto.UserRequest
	)

	if err := c.Bind(&request); err != nil {
		return c.JSON(400, err.Error())
	}

	if err := c.Validate(&request); err != nil {
		return c.JSON(400, err.Error())
	}

	user, err := h.Service.Update(ctx, request.ParseToModel(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(404, "Data Not Found")
		}
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, user)
}

func (h *UserController) Delete(c echo.Context) error {
	var (
		ctx   = c.Request().Context()
		id, _ = uuid.Parse(c.Param("id"))
	)

	if err := h.Service.Delete(ctx, id); err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, "Berhasil")
}
