package domain

import (
	"time"
)

type (
	RefreshToken     string
	JWTToken         string
	RefreshTokenData struct {
		Token     RefreshToken `json:"token,omitempty"`
		ExpiresAt time.Time    `json:"expiresAt"`
	}
	JWTBody struct {
		UserId UserId `json:"userId"`
	}
)
