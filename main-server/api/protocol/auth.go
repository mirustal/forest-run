package protocol

import (
	"forest-run/common/jwt"
	"forest-run/main-server/domain"
)

type (
	SignUpRequest struct {
		Username domain.Username `json:"username,omitempty"`
		Password domain.Password `json:"password,omitempty"`
	}
	SignUpResponse struct {
	}
)

type AuthDataResponse struct {
	RefreshToken jwt.RefreshToken `json:"refreshToken"`
	AuthToken    jwt.JWTTokenData `json:"authToken,omitempty"`
}

type (
	SignInRequest struct {
		SignUpRequest
	}
	SignInResponse struct {
		AuthDataResponse
	}
)

type (
	RefreshTokensRequest struct {
		AuthToken    jwt.JWTToken     `json:"authToken"`
		RefreshToken jwt.RefreshToken `json:"refreshToken"`
	}
	RefreshTokensResponse struct {
		AuthDataResponse
	}
)
