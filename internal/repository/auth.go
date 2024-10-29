package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository/interfaces"
	"gorm.io/gorm"
)

type authRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) interfaces.AuthRepository {
	return &authRepository{DB: db}
}

func (r *authRepository) GetTokenAuthByRefreshToken(ctx context.Context, token string) (data *model.TokenAuth, err error) {
	err = r.DB.Where("refresh_token = ?", token).First(&data).Error
	return
}

func (r *authRepository) GetTokenAuthByUserIDAndIP(ctx context.Context, userId, ip string) (data *model.TokenAuth, err error) {
	err = r.DB.Where("user_id = ? AND ip = ?", userId, ip).First(&data).Error
	return
}

func (r *authRepository) Store(ctx context.Context, payload *model.TokenAuth) error {
	fmt.Println("Store token")
	token, err := r.GetTokenAuthByUserIDAndIP(ctx, payload.UserID, payload.IP)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Error store token:", err)
			// Create a new data if not found
			if err := r.DB.Create(payload).Error; err != nil {
				return err
			}
			return nil
		}
		return err
	}

	fmt.Println("update token:", err)
	// update if exists user_id and ip
	if err := r.DB.Model(&model.TokenAuth{}).
		Where("user_id = ? AND ip = ?", token.UserID, token.IP).
		Update("refresh_token", payload.RefreshToken).Error; err != nil {
		return err
	}

	return nil
}

func (r *authRepository) Delete(ctx context.Context, refreshToken string) error {
	return r.DB.Where("refresh_token = ?", refreshToken).Delete(&model.TokenAuth{}).Error
}
