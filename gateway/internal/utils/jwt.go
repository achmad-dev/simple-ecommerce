package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTCustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type JwtHelper interface {
	GenerateToken(email string) (string, error)
	ValidateToken(tokenString string) (string, error)
}

type baseJwtHelper struct {
	SecretKey string
}

func NewJwtHelper(secret string) JwtHelper {
	return &baseJwtHelper{SecretKey: secret}
}

func (b *baseJwtHelper) GenerateToken(email string) (string, error) {
	claims := &JWTCustomClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(72 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(b.SecretKey)
}

func (b *baseJwtHelper) ValidateToken(tokenString string) (string, error) {
	claims := &JWTCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return b.SecretKey, nil
	})

	if _, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims.Email, nil
	}

	return "", err
}
