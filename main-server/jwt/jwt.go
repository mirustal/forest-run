package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"main-server/boot"
	"main-server/domain"
	"math/rand"
	"time"
)

type Provider interface {
	CreateToken(body domain.JWTBody) (domain.JWTToken, error)
	CreateRefreshToken() (domain.RefreshTokenData, error)
	Parse(token domain.JWTToken) (domain.JWTBody, error)
}

func NewProvider(cfg boot.JWTConfig) (Provider, error) {
	if cfg.JWTTokenLifeTime <= 0 {
		return nil, errors.New("JWTTokenLifeTime <= 0")
	}

	if len(cfg.SecureKey) == 0 {
		return nil, errors.New("length of secure key is 0")
	}

	if cfg.RefreshTokenLifeTime <= 0 {
		return nil, errors.New("RefreshTokenLifeTime <= 0")
	}

	return jwtProvider{
		JWTConfig: cfg,
	}, nil
}

type jwtProvider struct {
	boot.JWTConfig
}

type jwtClaims struct {
	domain.JWTBody
	jwt.RegisteredClaims
}

func (j jwtProvider) CreateToken(body domain.JWTBody) (token domain.JWTToken, err error) {
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, jwtClaims{
		JWTBody: body,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.JWTTokenLifeTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})
	s, err := t.SignedString(j.SecureKey)
	if err != nil {
		return token, err
	}

	return domain.JWTToken(s), err
}

func (j jwtProvider) Parse(token domain.JWTToken) (body domain.JWTBody, err error) {
	t, err := jwt.ParseWithClaims(string(token), &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SecureKey), nil
	})

	if err != nil {
		return body, err
	}

	if t.Valid == false {
		return body, errors.New("jwt token invalid")
	}

	if claims, ok := t.Claims.(*jwtClaims); ok {
		return claims.JWTBody, nil
	}

	return body, errors.New("unable to parse jwt claims")
}

func (j jwtProvider) CreateRefreshToken() (data domain.RefreshTokenData, err error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return data, err
	}

	data.Token = domain.RefreshToken(fmt.Sprintf("%x", b))
	data.ExpiresAt = time.Now().Add(j.RefreshTokenLifeTime)

	return data, err
}
