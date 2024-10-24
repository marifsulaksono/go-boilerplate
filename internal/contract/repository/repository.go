package repository

import (
	"context"

	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract/common"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository/interfaces"
)

type Contract struct {
	User interfaces.UserRepository
	Auth interfaces.AuthRepository
}

func NewRepository(ctx context.Context, common *common.Contract) (*Contract, error) {
	user := repository.NewUserRepository(common.DB)
	auth := repository.NewAuthRepository(common.DB)

	return &Contract{
		User: user,
		Auth: auth,
	}, nil
}
