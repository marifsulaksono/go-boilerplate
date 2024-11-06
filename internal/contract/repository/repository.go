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
	Role interfaces.RoleRepository
}

func NewRepository(ctx context.Context, common *common.Contract) (*Contract, error) {
	role := repository.NewRoleRepository(common.DB)
	user := repository.NewUserRepository(common.DB, common.Redis)
	auth := repository.NewAuthRepository(common.DB)

	return &Contract{
		User: user,
		Auth: auth,
		Role: role,
	}, nil
}
