package interfaces

import (
	"context"

	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
)

type AuthRepository interface {
	GetTokenAuthByRefreshToken(ctx context.Context, token string) (data *model.TokenAuth, err error)
	GetTokenAuthByUserIDAndIP(ctx context.Context, userId, ip string) (data *model.TokenAuth, err error)
	Store(ctx context.Context, payload *model.TokenAuth) error
	Delete(ctx context.Context, refreshToken string) error
}
