package interfaces

import (
	"context"

	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
)

type AuthService interface {
	Login(ctx context.Context, payload *model.Login, ip string) (model.LoginResponse, error)
	RefreshAccessToken(ctx context.Context, refreshToken string) (*model.LoginResponse, error)
	Logout(ctx context.Context, refreshToken string) error
}
