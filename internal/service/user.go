package service

import (
	"context"

	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract/repository"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository/interfaces"
)

type userService struct {
	Repository interfaces.UserRepository
}

type UserService interface {
	Get(ctx context.Context) string
}

func NewUserService(r *repository.Contract) UserService {
	return &userService{
		Repository: r.User,
	}
}

func (s *userService) Get(ctx context.Context) string {
	return s.Repository.Get(ctx)
}
