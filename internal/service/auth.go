package service

import (
	"context"
	"log"
	"net/http"

	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract/repository"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/helper"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/utils/response"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository/interfaces"
	sinterface "github.com/marifsulaksono/go-echo-boilerplate/internal/service/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	AuthRepository interfaces.AuthRepository
	UserRepository interfaces.UserRepository
}

func NewAuthService(r *repository.Contract) sinterface.AuthService {
	return &authService{
		AuthRepository: r.Auth,
		UserRepository: r.User,
	}
}

func (s *authService) Login(ctx context.Context, payload *model.Login, ip string) (model.LoginResponse, error) {
	user, err := s.UserRepository.GetByEmail(ctx, payload.Email)
	if err != nil {
		return model.LoginResponse{}, response.NewCustomError(http.StatusUnauthorized, "Informasi email atau password tidak sesuai", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		log.Printf("Error compare password: %v", err)
		return model.LoginResponse{}, response.NewCustomError(http.StatusUnauthorized, "Informasi email atau password tidak sesuai", err)
	}

	accessToken, expiredAt, err := helper.GenerateTokenJWT(user, false)
	if err != nil {
		log.Printf("Gagal generate access token: %v", err)
		return model.LoginResponse{}, response.NewCustomError(http.StatusInternalServerError, "Gagal menerbitkan token", nil)
	}

	refreshToken, _, err := helper.GenerateTokenJWT(user, true)
	if err != nil {
		log.Printf("Gagal generate refresh token: %v", err)
		return model.LoginResponse{}, response.NewCustomError(http.StatusInternalServerError, "Gagal menerbitkan token", nil)
	}

	if err := s.AuthRepository.Store(ctx, &model.TokenAuth{
		RefreshToken: refreshToken,
		UserID:       user.ID.String(),
		IP:           ip,
	}); err != nil {
		log.Printf("Gagal menyimpan token ke database: %v", err)
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
