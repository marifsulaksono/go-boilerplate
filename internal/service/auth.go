package service

import (
	"context"
	"errors"
	"log"

	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract/repository"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/helper"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	AuthRepository interfaces.AuthRepository
	UserRepository interfaces.UserRepository
}

type AuthService interface {
	Login(ctx context.Context, payload *model.Login) (model.LoginResponse, error)
}

func NewAuthService(r *repository.Contract) AuthService {
	return &authService{
		AuthRepository: r.Auth,
		UserRepository: r.User,
	}
}

func (s *authService) Login(ctx context.Context, payload *model.Login) (model.LoginResponse, error) {
	user, err := s.UserRepository.GetByEmail(ctx, payload.Email)
	if err != nil {
		return model.LoginResponse{}, errors.New("informasi email atau password tidak sesuai")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		log.Println("Error compare:", err)
		return model.LoginResponse{}, errors.New("informasi email atau password tidak sesuai")
	}

	accessToken, expiredAt, err := helper.GenerateTokenJWT(user, false)
	if err != nil {
		log.Printf("Gagal generate access token: %v", err)
		return model.LoginResponse{}, errors.New("gagal menerbitkan token")
	}

	refreshToken, _, err := helper.GenerateTokenJWT(user, true)
	if err != nil {
		log.Printf("Gagal generate refresh token: %v", err)
		return model.LoginResponse{}, errors.New("gagal menerbitkan token")
	}

	return model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Metadata: map[string]interface{}{
			"name":       user.Name,
			"email":      user.Email,
			"expired_at": expiredAt,
		},
	}, nil
}
