package jwt

import (
	"fmt"
	"forest-run/common"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestJWTTokens(t *testing.T) {
	provider := NewProvider(JWTConfig{
		SecureKey:            "qwerty",
		JWTTokenLifeTime:     1800000000000,
		RefreshTokenLifeTime: 2592000000000000,
	})

	token, err := provider.CreateToken(JWTBody{UserId: 1})
	assert.Nil(t, err)

	t.Logf("token: %v", token)

	body, err := provider.Parse(token.Token)
	assert.Nil(t, err)

	assert.Equal(t, common.UserId(1), body.UserId)

	splitted := strings.Split(string(token.Token), ".")
	assert.Len(t, splitted, 3)

	wrongToken := fmt.Sprintf("%v.%v.123", splitted[0], splitted[1])
	_, err = provider.Parse(JWTToken(wrongToken))
	assert.Error(t, err)

	wrongToken = fmt.Sprintf("%v.asdfds.%v", splitted[0], splitted[2])
	_, err = provider.Parse(JWTToken(wrongToken))
	assert.Error(t, err)
}

func TestExpiredJWTToken(t *testing.T) {
	provider := NewProvider(JWTConfig{
		SecureKey:            "qwerty",
		JWTTokenLifeTime:     -1800000000000,
		RefreshTokenLifeTime: 2592000000000000,
	})

	token, err := provider.CreateToken(JWTBody{UserId: 1})
	assert.Nil(t, err)

	t.Logf("token: %v", token)

	_, err = provider.Parse(token.Token)
	assert.Error(t, err)

	body, err := provider.ParseUnverified(token.Token)
	assert.Nil(t, err)
	assert.Equal(t, common.UserId(1), body.UserId)
}
