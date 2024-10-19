package repository

import (
	"context"

	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository/contract"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) contract.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) Get(ctx context.Context) string {
	return "BERHASIL!"
}
