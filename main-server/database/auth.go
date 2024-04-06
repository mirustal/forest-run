package database

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"main-server/domain"
)

func (p PgDbAdapter) StoreNewUser(username domain.Username, password domain.HashedPassword, ctx context.Context) error {
	t, err := p.dbPool.Begin(ctx)
	if err != nil {
		t.Rollback(ctx)
		return err
	}

	_, err = t.Exec(ctx, "INSERT INTO users (username, password) VALUES ($1, $2)", username, password)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return UsernameAlreadyTakenError{}
			}
		}
		return err
	}

	return t.Commit(ctx)
}

func (p PgDbAdapter) GetUserByUsername(username domain.Username, ctx context.Context) (user domain.User, err error) {
	t, err := p.dbPool.Begin(ctx)
	if err != nil {
		t.Rollback(ctx)
		return user, err
	}

	row := t.QueryRow(ctx, "SELECT id, username, password, refresh_token, refresh_token_expires_at FROM users WHERE username=$1", username)
	err = row.Scan(&user.Id, &user.Username, &user.HashedPassword, &user.RefreshTokenData.Token, &user.RefreshTokenData.ExpiresAt)
	if err != nil {
		return user, err
	}

	return user, t.Commit(ctx)
}

func (p PgDbAdapter) GetUserById(id domain.UserId, ctx context.Context) (user domain.User, err error) {
	t, err := p.dbPool.Begin(ctx)
	if err != nil {
		t.Rollback(ctx)
		return user, err
	}

	row := t.QueryRow(ctx, "SELECT id, username, password, refresh_token, refresh_token_expires_at FROM users WHERE id=$1", id)
	err = row.Scan(&user.Id, &user.Username, &user.HashedPassword, &user.RefreshTokenData.Token, &user.RefreshTokenData.ExpiresAt)
	if err != nil {
		return user, err
	}

	return user, t.Commit(ctx)
}

func (p PgDbAdapter) StoreUserRefreshToken(id domain.UserId, data domain.RefreshTokenData, ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (p PgDbAdapter) GetUserRefreshToken(id domain.UserId, ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
