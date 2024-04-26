package database

import (
	"context"
	"errors"
	"fmt"
	"forest-run/common"
	"forest-run/common/jwt"
	"forest-run/main-server/domain"
	"github.com/jackc/pgx/v5/pgconn"
)

type AuthRepo interface {
	StoreNewUser(username domain.Username, password domain.HashedPassword, ctx context.Context) error
	GetUserByUsername(username domain.Username, ctx context.Context) (domain.User, error)
	StoreUserRefreshToken(id common.UserId, data jwt.RefreshTokenData, ctx context.Context) error
	GetUserRefreshToken(id common.UserId, ctx context.Context) (data jwt.RefreshTokenData, err error)
}

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
	defer t.Rollback(ctx)
	if err != nil {
		return user, err
	}

	row := t.QueryRow(ctx, "SELECT id, username, password, refresh_token, refresh_token_expires_at FROM users WHERE username=$1", username)
	err = row.Scan(&user.Id, &user.Username, &user.HashedPassword, &user.RefreshTokenData.Token, &user.RefreshTokenData.ExpiresAt)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (p PgDbAdapter) GetUserById(id common.UserId, ctx context.Context) (user domain.User, err error) {
	t, err := p.dbPool.Begin(ctx)
	defer t.Rollback(ctx)
	if err != nil {
		return user, err
	}

	row := t.QueryRow(ctx, "SELECT id, username, password, refresh_token, refresh_token_expires_at FROM users WHERE id=$1", id)
	err = row.Scan(&user.Id, &user.Username, &user.HashedPassword, &user.RefreshTokenData.Token, &user.RefreshTokenData.ExpiresAt)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (p PgDbAdapter) StoreUserRefreshToken(id common.UserId, data jwt.RefreshTokenData, ctx context.Context) error {
	t, err := p.dbPool.Begin(ctx)
	if err != nil {
		defer t.Rollback(ctx)
		return err
	}

	c, err := t.Exec(ctx, "UPDATE users SET refresh_token=$1, refresh_token_expires_at=$2 WHERE id=$3", data.Token, data.ExpiresAt, id)
	if err != nil {
		defer t.Rollback(ctx)
		return nil
	}

	if c.RowsAffected() == 0 {
		defer t.Rollback(ctx)
		return errors.New(fmt.Sprintf("no rows were affected when storing refresh token for user %v", id))
	}

	return t.Commit(ctx)
}

func (p PgDbAdapter) GetUserRefreshToken(id common.UserId, ctx context.Context) (data jwt.RefreshTokenData, err error) {
	t, err := p.dbPool.Begin(ctx)
	defer t.Rollback(ctx)
	if err != nil {
		return data, err
	}

	row := t.QueryRow(ctx, "SELECT refresh_token, refresh_token_expires_at FROM users WHERE id=$1", id)
	err = row.Scan(&data.Token, &data.ExpiresAt)
	if err != nil {
		return data, err
	}

	return data, nil
}
