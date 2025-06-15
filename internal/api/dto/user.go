package dto

import (
	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
)

type (
	GetUserRequest struct {
		Page   int    `json:"page" query:"page" validate:"gte=1"`
		Limit  int    `json:"limit" query:"limit" validate:"gte=1"`
		Search string `json:"search" query:"search"`
	}

	UserRequest struct {
		Name     string    `json:"name" validate:"required"`
		Email    string    `json:"email" validate:"required,email"`
		Password string    `json:"password" validate:"required"`
		RoleID   uuid.UUID `json:"role_id" validate:"required"`
	}
)

func (u *GetUserRequest) ParseToModel() *model.UserRequest {
	return &model.UserRequest{
		Page:   u.Page,
		Limit:  u.Limit,
		Search: u.Search,
	}
}

func (u *UserRequest) ParseToModel() *model.User {
	return &model.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		RoleID:   u.RoleID,
	}
}
