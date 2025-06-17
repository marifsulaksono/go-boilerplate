package service

import (
	"context"

	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract/repository"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/service"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/service/interfaces"
)

type ServiceContract interface {
	GetUser() interfaces.UserService
	GetAuth() interfaces.AuthService
	GetRole() interfaces.RoleService
}

type Contract struct {
	User interfaces.UserService
	Auth interfaces.AuthService
	Role interfaces.RoleService
}

// NewService is used to prepare service dependency injection
func NewService(ctx context.Context, r repository.RepositoryContract) (ServiceContract, error) {
	user := service.NewUserService(r)
	auth := service.NewAuthService(r)
	role := service.NewRoleService(r)

	return &Contract{
		User: user,
		Auth: auth,
		Role: role,
	}, nil
}

func (s *Contract) GetUser() interfaces.UserService {
	return s.User
}

func (s *Contract) GetAuth() interfaces.AuthService {
	return s.Auth
}

func (s *Contract) GetRole() interfaces.RoleService {
	return s.Role
}
