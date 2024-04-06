package domain

import "time"

type (
	UserId           int
	Username         string
	Password         string
	RefreshToken     string
	RefreshTokenData struct {
		Token     RefreshToken `json:"token,omitempty"`
		ExpiresAt time.Time    `json:"expiresAt"`
	}
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}
