package protocol

type ErrorCode int

type ErrorResponse struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message,omitempty"`
}

const (
	CodeUserNameAlreadyTaken ErrorCode = 1
	CodeWrongPassword        ErrorCode = 2
)
