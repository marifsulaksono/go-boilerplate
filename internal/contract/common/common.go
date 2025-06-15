package common

import (
	"context"
	"log"

	"github.com/marifsulaksono/go-echo-boilerplate/internal/config"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/constants"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Contract struct {
	DB    *gorm.DB
	Redis *redis.Client
}

// NewCommon is used to prepare common dependency injection
func NewCommon(ctx context.Context) (*Contract, error) {
	db, err := config.Config.Database.ConnectDatabase(ctx, constants.DB_MYSQL)
	if err != nil {
		return nil, err
	}

	rediscli, err := config.Config.Redis.InitRedisClient()
	if err != nil {
		return nil, err
	}

	return &Contract{
		DB:    db,
		Redis: rediscli,
	}, nil
}

// GORM Migration
func (c *Contract) AutoMigrate() {
	if err := c.DB.AutoMigrate(
		&model.User{},
		&model.TokenAuth{},
	); err != nil {
		log.Fatalf("Error on migration database: %v", err)
	}
	log.Println("Migration successfully.....")
}

// Close all connection
func (c *Contract) Close() {
	c.Redis.Close()
	sqlDB, err := c.DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
}
