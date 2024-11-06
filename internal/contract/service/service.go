package service

import (
	"context"

	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract/repository"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/service"
)

type Contract struct {
	User service.UserService
	Auth service.AuthService
	Role service.RoleService
}

func NewService(ctx context.Context, r *repository.Contract) (*Contract, error) {
	user := service.NewUserService(r)
	auth := service.NewAuthService(r)
	role := service.NewRoleService(r)

	return &Contract{
		User: user,
		Auth: auth,
		Role: role,
	}, nil
}
