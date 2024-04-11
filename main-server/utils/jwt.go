package utils

import (
	"main-server/domain"
)

type JWTProvider interface {
	CreateToken(body domain.JWTBody) (string, error)
	ValidateToken(tokenString domain.JWTToken) error
	CreateRefreshToken() (string, error)
}

func NewJWTProvider(secureKey string) (JWTProvider, error) {
	return nil, nil
}
