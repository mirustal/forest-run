package domain

type (
	SignUpRequest struct {
		Username Username `json:"username,omitempty"`
		Password Password `json:"password,omitempty"`
	}
	SignUpResponse struct {
	}
)

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
