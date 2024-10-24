package config

import "github.com/golang-jwt/jwt/v5"

type (
	JWT struct {
		AccessSecret       string `json:"access_secret_key"`
		RefreshSecret      string `json:"refresh_secret_key"`
		AccessExpiryInSec  int    `json:"access_expiry_in_second"`
		RefreshExpiryInSec int    `json:"refresh_expiry_in_second"`
	}

	JWTClaim struct {
		ID    string `json:"id"`
		Email string `json:"email"`
		jwt.RegisteredClaims
	}
)
