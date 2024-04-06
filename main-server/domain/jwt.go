package domain

import "time"

type (
	RefreshToken     string
	JWTToken         string
	RefreshTokenData struct {
		Token     RefreshToken `json:"token,omitempty"`
		ExpiresAt time.Time    `json:"expiresAt"`
	}
	JWTBody struct {
		// TODO
	}
)

func NewRefreshToken(salt string) RefreshToken {
	panic("not implemented")
}

func NewJWTToken(body JWTBody, salt string) JWTToken {
	panic("not implemented")
}
