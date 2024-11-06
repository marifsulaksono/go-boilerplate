package repository

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/helper"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository/interfaces"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func NewUserRepository(db *gorm.DB, rdscli *redis.Client) interfaces.UserRepository {
	return &userRepository{
		DB:    db,
		Redis: rdscli,
	}
}

func (r *userRepository) Get(ctx context.Context) (data *[]model.User, err error) {
	const usersDataKey = "list-users-data"
	cachedUsers, err := r.Redis.Get(ctx, usersDataKey).Result()
	if err == nil {
		if err = json.Unmarshal([]byte(cachedUsers), &data); err != nil {
			log.Printf("Error unmarshaling user data from Redis: %v", err)
		} else {
			return data, nil
		}
	} else if err != redis.Nil {
		log.Printf("Error fetching user data from Redis: %v", err)
	}

	data = &[]model.User{}
	if err := r.DB.Joins("Role").Find(&data).Error; err != nil {
		return nil, err
	}

	err = helper.SetRedisJSONCache(ctx, r.Redis, usersDataKey, data, time.Duration(300)*time.Second)
	if err != nil {
		log.Printf("Error setting user data in Redis: %v", err)
	}

	return data, nil
}

func (r *userRepository) GetWithPagination(ctx context.Context, params *model.Pagination) (data *model.PaginationResponse, err error) {
	var users []model.User
	offset := (params.Page - 1) * params.Limit
	db := r.DB
	if params.Page > 0 {
		db = db.Offset(offset)
	}

	if params.Limit > 0 {
		db = db.Limit(params.Limit)
	}

	err = db.Joins("Role").Find(&users).Error
	if err != nil {
		return nil, err
	}

	var count int64
	err = r.DB.Model(&model.User{}).Where("deleted_at IS NULL").Count(&count).Error
	if err != nil {
		return nil, err
	}

	data = &model.PaginationResponse{
		List:         users,
		Page:         params.Page,
		Limit:        params.Limit,
		TotalPerPage: len(users),
		TotalData:    int(count),
	}

	return
}

func (r *userRepository) GetById(ctx context.Context, id uuid.UUID) (data *model.User, err error) {
	err = r.DB.Joins("Role").First(&data, id).Error
	return
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (data *model.User, err error) {
	err = r.DB.Where("email = ?", email).First(&data).Error
	return
}

func (r *userRepository) Create(ctx context.Context, payload *model.User) (string, error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok || userID == "" {
		payload.CreatedBy = ""
	} else {
		payload.CreatedBy = userID
	}

	err := r.DB.WithContext(ctx).Create(&payload).Clauses(clause.Returning{
		Columns: []clause.Column{
			{Name: "id"},
		},
	}).Error

	return payload.ID.String(), err
}

func (r *userRepository) Update(ctx context.Context, payload *model.User, id uuid.UUID) (string, error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok || userID == "" {
		payload.UpdatedBy = ""
	} else {
		payload.UpdatedBy = userID
	}
	err := r.DB.Model(&model.User{}).Where("id = ?", id).Updates(payload).Error
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.DB.WithContext(ctx).Where("id = ?", id).Delete(&model.User{}).Error
}
