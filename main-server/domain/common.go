package domain

type (
	UserId   int
	Username string
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}
