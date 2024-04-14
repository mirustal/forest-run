package auth

import (
	"errors"
	"main-server/domain"
)

func Validate(r domain.SignUpRequest) error {
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
