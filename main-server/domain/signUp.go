package domain

type (
	SignUpRequest struct {
		Login    Login    `json:"login,omitempty"`
		Password Password `json:"password,omitempty"`
	}
	SignUpResponse struct {
		AuthToken AuthToken `json:"success,omitempty"`
	}
)
