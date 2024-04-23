package domain

import "forest-run/common"

type (
	Username string
)

type User struct {
	Id               common.UserId
	Username         Username
	HashedPassword   HashedPassword
	RefreshTokenData RefreshTokenData
}
