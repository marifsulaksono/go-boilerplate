package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository/interfaces"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) Get(ctx context.Context) (data *[]model.User, err error) {
	err = r.DB.Find(&data).Error
	return
}

func (r *userRepository) GetById(ctx context.Context, id uuid.UUID) (data *model.User, err error) {
	err = r.DB.First(&data, id).Error
	return
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (data *model.User, err error) {
	err = r.DB.Where("email = ?", email).First(&data).Error
	return
}

func (r *userRepository) Create(ctx context.Context, payload *model.User) (data *model.UserResponse, err error) {
	data = new(model.UserResponse)
	err = r.DB.WithContext(ctx).Create(&payload).Clauses(clause.Returning{
		Columns: []clause.Column{
			{Name: "id"},
			{Name: "name"},
			{Name: "email"},
			{Name: "created_at"},
			{Name: "updated_at"},
		},
	}).Error

	if err != nil {
		return nil, err
	}

	data.ID = payload.ID
	data.Name = payload.Name
	data.Email = payload.Email
	data.CreatedAt = payload.CreatedAt
	data.UpdatedAt = payload.UpdatedAt

	return
}

func (r *userRepository) Update(ctx context.Context, payload *model.User, id uuid.UUID) (data *model.UserResponse, err error) {
	data = new(model.UserResponse)

	err = r.DB.Model(&model.User{}).
		Where("id = ?", id).
		Updates(payload).Error

	if err != nil {
		return nil, err
	}

	data.ID = payload.ID
	data.Name = payload.Name
	data.Email = payload.Email
	data.CreatedAt = payload.CreatedAt
	data.UpdatedAt = payload.UpdatedAt

	return data, nil
}

func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.DB.WithContext(ctx).Where("id = ?", id).Delete(&model.User{}).Error
}
