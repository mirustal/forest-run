package domain

type User struct {
	Id               UserId
	Username         Username
	HashedPassword   HashedPassword
	RefreshTokenData RefreshTokenData
}
