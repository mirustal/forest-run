package domain

import (
	"time"
)

type (
	RefreshToken string
	JWTToken     string
	JWTTokenData struct {
		Token     JWTToken  `json:"token,omitempty"`
		ExpiresAt time.Time `json:"expiresAt,omitempty"`
	}
	RefreshTokenData struct {
		Token     *RefreshToken `json:"token,omitempty"`
		ExpiresAt *time.Time    `json:"expiresAt,omitempty"`
	}
	JWTBody struct {
		UserId UserId `json:"userId"`
	}
)
