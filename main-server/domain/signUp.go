package domain

type (
	SignUpRequest struct {
		Username Username `json:"login,omitempty"`
		Password Password `json:"password,omitempty"`
	}
	SignUpResponse struct {
	}
)
