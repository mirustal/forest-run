package domain

import (
	"errors"
)

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
