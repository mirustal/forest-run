package domain

import "golang.org/x/crypto/bcrypt"

type (
	Password       string
	HashedPassword string
)

const maxBcryptPasswordLength = 72

func (p Password) Hash() HashedPassword {
	b := []byte(p)
	if len(b) > maxBcryptPasswordLength {
		b = b[0 : maxBcryptPasswordLength-1]
	}
	bytes, _ := bcrypt.GenerateFromPassword(b, 14)
	return HashedPassword(bytes)
}

func (p Password) Matches(hash HashedPassword) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(p))
	return err == nil
}
