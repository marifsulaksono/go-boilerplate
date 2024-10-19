package repository

import (
	"context"

	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract/common"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository/contract"
)

type Contract struct {
	User contract.UserRepository
}

func NewRepository(ctx context.Context, common *common.Contract) (*Contract, error) {
	user := repository.NewUserRepository(common.DB)

	return &Contract{
		User: user,
	}, nil
}
