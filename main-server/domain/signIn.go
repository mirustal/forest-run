package domain

type (
	SignInRequest struct {
		SignUpRequest
	}
	SignInResponse struct {
		AuthDataResponse
	}
	AuthDataResponse struct {
		RefreshToken RefreshTokenData `json:"refreshToken"`
		AuthToken    JWTToken         `json:"authToken,omitempty"`
	}
)
