package service

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/constants"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract/repository"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/helper"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/utils/response"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository/interfaces"
	sinterface "github.com/marifsulaksono/go-echo-boilerplate/internal/service/interfaces"
	"gorm.io/gorm"
)

type authService struct {
	AuthRepository interfaces.AuthRepository
	UserRepository interfaces.UserRepository
	RoleRepository interfaces.RoleRepository
}

func NewAuthService(r *repository.Contract) sinterface.AuthService {
	return &authService{
		AuthRepository: r.Auth,
		UserRepository: r.User,
		RoleRepository: r.Role,
	}
}

func (s *authService) Register(ctx context.Context, payload *model.Register) error {
	user, err := s.UserRepository.GetByEmail(ctx, payload.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return response.NewCustomError(http.StatusInternalServerError, "Terjadi kesalahan pada server", err)
	}

	if user.ID != uuid.Nil {
		return response.NewCustomError(http.StatusBadRequest, "Email sudah terdaftar", nil)
	}

	password, err := helper.GenerateHashedPassword(payload.Password)
	if err != nil {
		return response.NewCustomError(http.StatusInternalServerError, "Terjadi kesalahan pada server", err)
	}

	role, err := s.RoleRepository.GetByName(ctx, constants.DEFAULT_ROLE)
	if err != nil {
		return response.NewCustomError(http.StatusInternalServerError, "Terjadi kesalahan pada server", err)
	}

	_, err = s.UserRepository.Create(ctx, &model.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: password,
		RoleID:   role.ID,
	})
	if err != nil {
		return response.NewCustomError(http.StatusInternalServerError, "Terjadi kesalahan pada server", err)
	}

	return nil
}

func (s *authService) Login(ctx context.Context, payload *model.Login) (model.LoginResponse, error) {
	user, err := s.UserRepository.GetByEmail(ctx, payload.Email)
	if err != nil {
		return model.LoginResponse{}, response.NewCustomError(http.StatusUnauthorized, "Informasi email atau password tidak sesuai", err)
	}

	err = helper.CompareHashedPassword(payload.Password, user.Password)
	if err != nil {
		return model.LoginResponse{}, response.NewCustomError(http.StatusUnauthorized, "Informasi email atau password tidak sesuai", err)
	}

	accessToken, expiredAt, err := helper.GenerateTokenJWT(user, false)
	if err != nil {
		return model.LoginResponse{}, response.NewCustomError(http.StatusInternalServerError, "Gagal menerbitkan access token", nil)
	}

	refreshToken, _, err := helper.GenerateTokenJWT(user, true)
	if err != nil {
		return model.LoginResponse{}, response.NewCustomError(http.StatusInternalServerError, "Gagal menerbitkan refresh token", nil)
	}

	if err := s.AuthRepository.Store(ctx, &model.TokenAuth{
		RefreshToken: refreshToken,
		UserID:       user.ID.String(),
		IP:           payload.IP,
	}); err != nil {
		return model.LoginResponse{}, response.NewCustomError(http.StatusInternalServerError, "Gagal ketika menyimpan token ke database", nil)
	}

	return model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Metadata: model.MetadataLoginResponse{
			Name:      user.Name,
			Email:     user.Email,
			ExpiredAt: *expiredAt,
		},
	}, nil
}

func (s *authService) RefreshAccessToken(ctx context.Context, refreshToken string) (*model.LoginResponse, error) {
	token, err := s.AuthRepository.GetTokenAuthByRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	user, err := helper.VerifyTokenJWT(token.RefreshToken, true)
	if err != nil {
		log.Printf("Gagal memverifikasi refresh token: %v", err)
		return nil, response.NewCustomError(http.StatusInternalServerError, "Gagal memverifikasi token", nil)
	}

	accessToken, expiredAt, err := helper.GenerateTokenJWT(user, false)
	if err != nil {
		log.Printf("Gagal generate access token: %v", err)
		return nil, response.NewCustomError(http.StatusInternalServerError, "Gagal menerbitkan token", nil)
	}

	return &model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Metadata: model.MetadataLoginResponse{
			Name:      user.Name,
			Email:     user.Email,
			ExpiredAt: *expiredAt,
		},
	}, nil
}

func (s *authService) Logout(ctx context.Context, refreshToken string) error {
	return s.AuthRepository.Delete(ctx, refreshToken)
}
