package database

import (
	"context"
	"main-server/domain"
)

func (p PgDbAdapter) StoreNewUser(username domain.Username, password domain.Password, ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (p PgDbAdapter) StoreUserRefreshToken(id domain.UserId, data domain.RefreshTokenData, ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (p PgDbAdapter) GetUserRefreshToken(id domain.UserId, ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
