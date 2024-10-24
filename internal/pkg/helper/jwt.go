package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/config"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
)

func GenerateTokenJWT(user *model.User, isRefresh bool) (string, *time.Time, error) {
	var (
		expiredInSecond int
		secretKey       string
	)

	if isRefresh {
		expiredInSecond = config.Config.JWT.RefreshExpiryInSec
		secretKey = config.Config.JWT.RefreshSecret
	} else {
		expiredInSecond = config.Config.JWT.AccessExpiryInSec
		secretKey = config.Config.JWT.AccessSecret
	}

	expiredAt := time.Now().Add(time.Second * time.Duration(expiredInSecond))
	claims := &config.JWTClaim{
		ID:    user.ID.String(),
		Email: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "venturo",
			ExpiresAt: jwt.NewNumericDate(expiredAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Declare the token with the HS256 algorithm used for signing, and the claims.
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
	if err != nil {
		return "", nil, err
	}

	return tokenString, &expiredAt, nil
}

func VerifyTokenJWT(tokenString string, isRefresh bool) (*model.User, error) {
	var (
		secretKey string
		user      = new(model.User)
	)

	if isRefresh {
		secretKey = config.Config.JWT.RefreshSecret
	} else {
		secretKey = config.Config.JWT.AccessSecret
	}

	token, err := jwt.ParseWithClaims(tokenString, &config.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	// extract user claims if the token is valid
	if claims, ok := token.Claims.(*config.JWTClaim); ok && token.Valid {

		user.ID, _ = uuid.Parse(claims.ID)
		user.Name = claims.Email

		return user, nil
	}

	return nil, err
}
