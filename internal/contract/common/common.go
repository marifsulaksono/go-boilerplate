package common

import (
	"context"
	"log"

	"github.com/marifsulaksono/go-echo-boilerplate/internal/config"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/constants"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"gorm.io/gorm"
)

type Contract struct {
	DB *gorm.DB
}

func NewCommon(ctx context.Context) (*Contract, error) {
	db, err := config.Config.Database.ConnectDatabase(ctx, constants.DB_POSTGRESQL)
	if err != nil {
		return nil, err
	}

	return &Contract{
		DB: db,
	}, nil
}

func (c *Contract) AutoMigrate() {
	if err := c.DB.AutoMigrate(
		&model.User{},
		&model.TokenAuth{},
	); err != nil {
		log.Fatalf("Error on migration database: %v", err)
	}
	log.Println("Migration successfully.....")
}
