package domain

import "time"

type AuthToken string

type AuthTokenData struct {
	Token     AuthToken
	CreatedAt time.Time
}
