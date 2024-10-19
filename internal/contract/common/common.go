package common

import (
	"context"

	"github.com/marifsulaksono/go-echo-boilerplate/internal/config"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/constants"
	"gorm.io/gorm"
)

type Contract struct {
	DB *gorm.DB
}

func NewCommon(ctx context.Context) (*Contract, error) {
	db, err := config.Config.Database.ConnectDatabase(ctx, constants.DB_MYSQL)
	if err != nil {
		return nil, err
	}

	return &Contract{
		DB: db,
	}, nil
}
