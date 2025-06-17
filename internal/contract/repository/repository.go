package repository

import (
	"context"

	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract/common"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository/interfaces"
)

type RepositoryContract interface {
	GetUser() interfaces.UserRepository
	GetAuth() interfaces.AuthRepository
	GetRole() interfaces.RoleRepository
}

type Contract struct {
	User interfaces.UserRepository
	Auth interfaces.AuthRepository
	Role interfaces.RoleRepository
}

// NewRepository is used to prepare repository dependency injection
func NewRepository(ctx context.Context, common *common.Contract) (RepositoryContract, error) {
	role := repository.NewRoleRepository(common.DB)
	user := repository.NewUserRepository(common.DB, common.Redis)
	auth := repository.NewAuthRepository(common.DB)

	return &Contract{
		User: user,
		Auth: auth,
		Role: role,
	}, nil
}

func (r *Contract) GetUser() interfaces.UserRepository {
	return r.User
}

func (r *Contract) GetAuth() interfaces.AuthRepository {
	return r.Auth
}

func (r *Contract) GetRole() interfaces.RoleRepository {
	return r.Role
}
