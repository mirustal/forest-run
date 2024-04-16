package auth

import (
	"golang.org/x/crypto/bcrypt"
	"main-server/domain"
)

const maxBcryptPasswordLength = 72

func hash(p domain.Password) domain.HashedPassword {
	b := []byte(p)
	if len(b) > maxBcryptPasswordLength {
		b = b[0 : maxBcryptPasswordLength-1]
	}
	bytes, _ := bcrypt.GenerateFromPassword(b, 14)
	return domain.HashedPassword(bytes)
}

func matches(p domain.Password, hash domain.HashedPassword) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(p))
	return err == nil
}
