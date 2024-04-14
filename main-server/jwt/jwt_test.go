package jwt

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"main-server/boot"
	"main-server/domain"
	"strings"
	"testing"
)

func TestJWTTokens(t *testing.T) {
	provider := NewProvider(boot.JWTConfig{
		SecureKey:            "qwerty",
		JWTTokenLifeTime:     1800000000000,
		RefreshTokenLifeTime: 2592000000000000,
	})

	token, err := provider.CreateToken(domain.JWTBody{UserId: 1})
	assert.Nil(t, err)

	t.Logf("token: %v", token)

	body, err := provider.Parse(token.Token)
	assert.Nil(t, err)

	assert.Equal(t, domain.UserId(1), body.UserId)

	splitted := strings.Split(string(token.Token), ".")
	assert.Len(t, splitted, 3)

	wrongToken := fmt.Sprintf("%v.%v.123", splitted[0], splitted[1])
	_, err = provider.Parse(domain.JWTToken(wrongToken))
	assert.Error(t, err)

	wrongToken = fmt.Sprintf("%v.asdfds.%v", splitted[0], splitted[2])
	_, err = provider.Parse(domain.JWTToken(wrongToken))
	assert.Error(t, err)
}

func TestExpiredJWTToken(t *testing.T) {
	provider := NewProvider(boot.JWTConfig{
		SecureKey:            "qwerty",
		JWTTokenLifeTime:     -1800000000000,
		RefreshTokenLifeTime: 2592000000000000,
	})

	token, err := provider.CreateToken(domain.JWTBody{UserId: 1})
	assert.Nil(t, err)

	t.Logf("token: %v", token)

	_, err = provider.Parse(token.Token)
	assert.Error(t, err)
}
