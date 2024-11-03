package model

import "time"

type (
	Login struct {
		GrantType string `json:"grant_type"`
		Username  string `json:"username"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	LoginResponse struct {
		AccessToken  string                `json:"access_token"`
		RefreshToken string                `json:"refresh_token"`
		Metadata     MetadataLoginResponse `json:"metadata"`
	}

	MetadataLoginResponse struct {
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		ExpiredAt time.Time `json:"expired_at"`
	}
)

type (
	TokenAuth struct {
		RefreshToken string `json:"refresh_token" gorm:"not null"`
		UserID       string `json:"user_id" gorm:"not null"`
		IP           string `json:"ip" gorm:"not null;varchar(128)"`
	}
)
