package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract/repository"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/helper"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/utils/response"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository/interfaces"
	sinterface "github.com/marifsulaksono/go-echo-boilerplate/internal/service/interfaces"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userService struct {
	UserRepository interfaces.UserRepository
}

func NewUserService(r *repository.Contract) sinterface.UserService {
	return &userService{
		UserRepository: r.User,
	}
}

func (s *userService) Get(ctx context.Context) (*[]model.User, error) {
	return s.UserRepository.Get(ctx)
}

func (s *userService) GetWithPagination(ctx context.Context, params *model.Pagination) (data *model.PaginationResponse, err error) {
	return s.UserRepository.GetWithPagination(ctx, params)
}

func (s *userService) GetById(ctx context.Context, id uuid.UUID) (data *model.User, err error) {
	data, err = s.UserRepository.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NewCustomError(http.StatusNotFound, "Data tidak ditemukan", nil)
		}
		return nil, response.NewCustomError(http.StatusInternalServerError, "Terjadi kesalahan pada server", err)
	}

	return
}

func (s *userService) GetByEmail(ctx context.Context, email string) (data *model.User, err error) {
	data, err = s.UserRepository.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NewCustomError(http.StatusNotFound, "Data tidak ditemukan", nil)
		}
		return nil, response.NewCustomError(http.StatusInternalServerError, "Terjadi kesalahan pada server", err)
	}

	return
}

func (s *userService) Create(ctx context.Context, payload *model.User) (id string, err error) {
	payload.Password, err = helper.GenerateHashedPassword(payload.Password)
	if err != nil {
		return "", err
	}
	return s.UserRepository.Create(ctx, payload)
}

func (s *userService) Update(ctx context.Context, payload *model.User, id uuid.UUID) (string, error) {
	_, err := s.UserRepository.GetById(ctx, id)
	if err != nil {
		return "", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	payload.Password = string(hashedPassword)
	return s.UserRepository.Update(ctx, payload, id)
}

func (s *userService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.UserRepository.Delete(ctx, id)
}
