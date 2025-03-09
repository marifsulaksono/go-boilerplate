package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
)

type UserService interface {
	Get(ctx context.Context) (*[]model.User, error)
	GetWithPagination(ctx context.Context, params *model.Pagination) (data *model.PaginationResponse, err error)
	GetById(ctx context.Context, id uuid.UUID) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Create(ctx context.Context, payload *model.User) (string, error)
	Update(ctx context.Context, payload *model.User, id uuid.UUID) (string, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
