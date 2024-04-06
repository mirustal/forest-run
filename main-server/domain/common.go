package domain

import "time"

type (
	UserId        int
	Login         string
	Password      string
	AuthToken     string
	AuthTokenData struct {
		Token     AuthToken `json:"token,omitempty"`
		CreatedAt time.Time `json:"createdAt"`
	}
)

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
}
