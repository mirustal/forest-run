package database

import (
	"context"
	"errors"
	"main-server/domain"
)

func (p PgDbAdapter) DropAuthTokens(ctx context.Context, id domain.UserId) error {
	t, err := p.dbPool.Begin(ctx)
	if err != nil {
		t.Rollback(ctx)
		return err
	}

	_, err = t.Exec(ctx, "DELETE FROM auth_tokens WHERE user_id=$1", id)
	if err != nil {
		t.Rollback(ctx)
		return err
	}

	return t.Commit(ctx)
}

func (p PgDbAdapter) GetLatestToken(ctx context.Context, id domain.UserId) (*domain.AuthTokenData, error) {
	t, err := p.dbPool.Begin(ctx)
	if err != nil {
		t.Rollback(ctx)
		return nil, err
	}

	row := t.QueryRow(ctx, "SELECT token, created_at FROM auth_tokens WHERE user_id=$1 ORDER BY created_at DESC LIMIT 1", id)
	token := &domain.AuthTokenData{}
	err = row.Scan(token)
	if err != nil {
		t.Rollback(ctx)
		return nil, err
	}

	return token, t.Commit(ctx)
}

func (p PgDbAdapter) StoreToken(ctx context.Context, id domain.UserId, token domain.AuthToken) error {
	t, err := p.dbPool.Begin(ctx)
	if err != nil {
		t.Rollback(ctx)
		return err
	}

	c, err := t.Exec(ctx, "INSERT INTO auth_tokens (user_id, token) VALUES ($1, $2)", id, token)
	if err != nil {
		t.Rollback(ctx)
		return err
	}

	if c.RowsAffected() == 0 {
		t.Rollback(ctx)
		return errors.New("token stored without any rows affected")
	}

	return t.Commit(ctx)
}
