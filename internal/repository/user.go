package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
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

func (r *userRepository) Get(ctx context.Context, params *model.UserRequest) (data []model.User, total int64, err error) {
	var (
		offset = (params.Page - 1) * params.Limit
	)

	db := r.DB
	if params.Search != "" {
		db = db.Where("users.name ILIKE ? OR users.email ILIKE ?", "%"+params.Search+"%", "%"+params.Search+"%")
	}

	err = db.Joins("Role").Offset(offset).Limit(params.Limit).Find(&data).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.DB.Model(&model.User{}).Where("deleted_at IS NULL").Count(&total).Error
	if err != nil {
		return nil, 0, err
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
