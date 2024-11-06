package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract/repository"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository/interfaces"
)

type roleService struct {
	roleRepo interfaces.RoleRepository
}

type RoleService interface {
	Get(ctx context.Context) (data *[]model.Role, err error)
	GetById(ctx context.Context, id uuid.UUID) (data *model.Role, err error)
	Create(ctx context.Context, payload *model.Role) (string, error)
	Update(ctx context.Context, payload *model.Role, id uuid.UUID) (string, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

func NewRoleService(r *repository.Contract) RoleService {
	return &roleService{
		roleRepo: r.Role,
	}
}

func (s *roleService) Get(ctx context.Context) (data *[]model.Role, err error) {
	return s.roleRepo.Get(ctx)
}

func (s *roleService) GetById(ctx context.Context, id uuid.UUID) (data *model.Role, err error) {
	return s.roleRepo.GetById(ctx, id)
}

func (s *roleService) Create(ctx context.Context, payload *model.Role) (string, error) {
	return s.roleRepo.Create(ctx, payload)
}

func (s *roleService) Update(ctx context.Context, payload *model.Role, id uuid.UUID) (string, error) {
	return s.roleRepo.Update(ctx, payload, id)
}

func (s *roleService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.roleRepo.Delete(ctx, id)
}
