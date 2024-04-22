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
	CreateToken(body domain.JWTBody) (domain.JWTTokenData, error)
	CreateRefreshToken() (domain.RefreshTokenData, error)
	Parse(token domain.JWTToken) (domain.JWTBody, error)
	ParseUnverified(token domain.JWTToken) (body domain.JWTBody, err error)
}

func NewProvider(cfg boot.JWTConfig) Provider {
	return jwtProvider{
		JWTConfig: cfg,
	}
}

type jwtProvider struct {
	boot.JWTConfig
}

type jwtClaims struct {
	domain.JWTBody
	jwt.RegisteredClaims
}

func (j jwtProvider) CreateToken(body domain.JWTBody) (token domain.JWTTokenData, err error) {
	expTime := time.Now().Add(j.JWTTokenLifeTime)

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims{
		JWTBody: body,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})
	s, err := t.SignedString([]byte(j.SecureKey))
	if err != nil {
		return token, err
	}

	return domain.JWTTokenData{
		Token:     domain.JWTToken(s),
		ExpiresAt: expTime,
	}, err
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

func (j jwtProvider) ParseUnverified(token domain.JWTToken) (body domain.JWTBody, err error) {
	t, _ := jwt.ParseWithClaims(string(token), &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SecureKey), nil
	}, jwt.WithoutClaimsValidation())

	if t == nil {
		return body, errors.New("unable to parse jwt token")
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

	t := domain.RefreshToken(fmt.Sprintf("%x", b))
	exp := time.Now().Add(j.RefreshTokenLifeTime)

	data = domain.RefreshTokenData{
		Token:     &t,
		ExpiresAt: &exp,
	}

	return data, err
}
