package dto

import "github.com/marifsulaksono/go-echo-boilerplate/internal/model"

type (
	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
)

func (l *LoginRequest) ParseToModel() *model.Login {
	return &model.Login{
		Email:    l.Email,
		Password: l.Password,
	}
}
