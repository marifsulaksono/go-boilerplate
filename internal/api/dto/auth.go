package dto

import "github.com/marifsulaksono/go-echo-boilerplate/internal/model"

type (
	RegisterRequest struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	RefreshAccessTokenRequest struct {
		RefreshToken string `json:"refresh_token" validate:"required"`
	}
)

func (d *LoginRequest) ParseToModel(ip string) *model.Login {
	return &model.Login{
		Email:    d.Email,
		Password: d.Password,
		IP:       ip,
	}
}

func (d *RegisterRequest) ParseToModel() *model.Register {
	return &model.Register{
		Name:     d.Name,
		Email:    d.Email,
		Password: d.Password,
	}
}
