package handler

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	jwt.RegisteredClaims
	ID	uuid.UUID	`json:"id"`
}

func Generate(id uuid.UUID, secret string) (string, error) {
	now := time.Now()

	claims := Claims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(
				now.Add(24 * time.Hour),
			),
		},
	}

	jwtToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenString, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}