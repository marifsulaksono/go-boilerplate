package repository

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/utils/response"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository/interfaces"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type roleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) interfaces.RoleRepository {
	return &roleRepository{
		DB: db,
	}
}

func (r *roleRepository) GetWithPagination(ctx context.Context, params *model.Pagination) (data *model.PaginationResponse, err error) {
	var roles []model.Role
	offset := (params.Page - 1) * params.Limit
	db := r.DB
	if params.Page > 0 {
		db = db.Offset(offset)
	}

	if params.Limit > 0 {
		db = db.Limit(params.Limit)
	}

	err = db.Find(&roles).Error
	if err != nil {
		return nil, err
	}

	var count int64
	err = r.DB.Model(&model.Role{}).Where("deleted_at IS NULL").Count(&count).Error
	if err != nil {
		return nil, err
	}

	data = &model.PaginationResponse{
		List:         roles,
		Page:         params.Page,
		Limit:        params.Limit,
		TotalPerPage: len(roles),
		TotalData:    int(count),
	}
	return data, nil
}

func (r *roleRepository) Get(ctx context.Context) (data *[]model.Role, err error) {
	data = &[]model.Role{}
	if err := r.DB.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *roleRepository) GetById(ctx context.Context, id uuid.UUID) (data *model.Role, err error) {
	err = r.DB.First(&data, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NewCustomError(http.StatusNotFound, "Data tidak ditemukan", nil)
		}
		return nil, response.NewCustomError(http.StatusInternalServerError, "Terjadi kesalahan pada server", err)
	}

	return
}

func (r *roleRepository) Create(ctx context.Context, payload *model.Role) (string, error) {
	err := r.DB.WithContext(ctx).Create(&payload).Clauses(clause.Returning{
		Columns: []clause.Column{
			{Name: "id"},
		},
	}).Error

	return payload.ID.String(), err
}

func (r *roleRepository) Update(ctx context.Context, payload *model.Role, id uuid.UUID) (string, error) {
	err := r.DB.Model(&model.Role{}).Where("id = ?", id).Updates(payload).Error
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *roleRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.DB.Model(&model.Role{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error
}
