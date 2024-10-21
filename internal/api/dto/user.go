package dto

import "github.com/marifsulaksono/go-echo-boilerplate/internal/model"

type (
	UserRequest struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
)

func (u *UserRequest) ParseToModel() *model.User {
	return &model.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}
