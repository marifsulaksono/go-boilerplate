package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (m *UserRepository) Get(ctx context.Context) (*[]model.User, error) {
	args := m.Called(ctx)
	return args.Get(0).(*[]model.User), args.Error(1)
}

func (m *UserRepository) GetWithPagination(ctx context.Context, params *model.Pagination) (*model.PaginationResponse, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(*model.PaginationResponse), args.Error(1)
}

func (m *UserRepository) GetById(ctx context.Context, id uuid.UUID) (*model.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *UserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *UserRepository) Create(ctx context.Context, payload *model.User) (string, error) {
	args := m.Called(ctx, payload)
	return args.String(0), args.Error(1)
}

func (m *UserRepository) Update(ctx context.Context, payload *model.User, id uuid.UUID) (string, error) {
	args := m.Called(ctx, payload, id)
	return args.String(0), args.Error(1)
}

func (m *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
