package domain

import (
	"forest-run/common"
	"forest-run/common/jwt"
)

type (
	Username string
)

type User struct {
	Id               common.UserId
	Username         Username
	HashedPassword   HashedPassword
	RefreshTokenData jwt.RefreshTokenData
}
