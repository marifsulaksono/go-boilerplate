package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/config"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
)

func GenerateTokenJWT(user *model.User, isRefresh bool) (string, *time.Time, error) {
	var (
		expiredInSecond int
	)

	if isRefresh {
		expiredInSecond = config.Config.JWT.RefreshExpiryInSec
	} else {
		expiredInSecond = config.Config.JWT.AccessExpiryInSec
	}

	privateKeyBytes, err := os.ReadFile(config.Config.JWT.PrivateKeyPathFile)
	if err != nil {
		return "", nil, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return "", nil, err
	}

	expiredAt := time.Now().Add(time.Second * time.Duration(expiredInSecond))
	claims := &config.JWTClaim{
		ID:    user.ID.String(),
		Email: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "muhammad-arif-sulaksono",
			ExpiresAt: jwt.NewNumericDate(expiredAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Declare the token with the HS256 algorithm used for signing, and the claims.
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodRS512, claims).SignedString(privateKey)
	if err != nil {
		return "", nil, err
	}

	return tokenString, &expiredAt, nil
}

func VerifyTokenJWT(tokenString string, isRefresh bool) (*model.User, error) {
	var (
		user = new(model.User)
	)

	publicKeyBytes, err := os.ReadFile(config.Config.JWT.PublicKeyPathFile)
	if err != nil {
		return nil, err
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(tokenString, &config.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
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
