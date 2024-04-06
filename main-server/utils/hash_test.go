package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	passwords := []string{
		"secret", "qwerty123", "hhhaaa", "cdxs",
	}

	for _, password := range passwords {
		h, err := HashPassword(password)
		assert.Nil(t, err)
		assert.True(t, CheckPasswordHash(password, h))
	}
}
