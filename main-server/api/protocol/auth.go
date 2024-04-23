package protocol

import "forest-run/main-server/domain"

type (
	SignUpRequest struct {
		Username domain.Username `json:"username,omitempty"`
		Password domain.Password `json:"password,omitempty"`
	}
	SignUpResponse struct {
	}
)

type AuthDataResponse struct {
	RefreshToken domain.RefreshToken `json:"refreshToken"`
	AuthToken    domain.JWTTokenData `json:"authToken,omitempty"`
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
		AuthToken    domain.JWTToken     `json:"authToken"`
		RefreshToken domain.RefreshToken `json:"refreshToken"`
	}
	RefreshTokensResponse struct {
		AuthDataResponse
	}
)
