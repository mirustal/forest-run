package domain

import "errors"

type (
	SignUpRequest struct {
		Username Username `json:"username,omitempty"`
		Password Password `json:"password,omitempty"`
	}
	SignUpResponse struct {
	}
)

func (r SignUpRequest) Validate() error {
	if len(r.Username) == 0 {
		return errors.New("empty login")
	}

	if len(r.Username) > 64 {
		return errors.New("too long login")
	}

	if len(r.Password) == 0 {
		return errors.New("empty password")
	}

	if len(r.Password) > 64 {
		return errors.New("too long password")
	}

	return nil
}

type AuthDataResponse struct {
	RefreshToken RefreshToken `json:"refreshToken"`
	AuthToken    JWTTokenData `json:"authToken,omitempty"`
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
		RefreshToken RefreshToken `json:"refreshToken"`
	}
	RefreshTokensResponse struct {
		AuthDataResponse
	}
)
