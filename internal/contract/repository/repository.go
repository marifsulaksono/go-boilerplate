package repository

import (
	"context"

	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract/common"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository/interfaces"
)

type Contract struct {
	User interfaces.UserRepository
}

func NewRepository(ctx context.Context, common *common.Contract) (*Contract, error) {
	user := repository.NewUserRepository(common.DB)

	return &Contract{
		User: user,
	}, nil
}
