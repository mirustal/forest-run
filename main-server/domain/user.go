package domain

type (
	UserId   int
	Username string
)

type User struct {
	Id               UserId
	Username         Username
	HashedPassword   HashedPassword
	RefreshTokenData RefreshTokenData
}
