package database

import (
	"context"
	"main-server/domain"
)

type SubscriptionsRepo interface {
	Subscribe(subscriber domain.UserId, receiver domain.UserId, ctx context.Context) (bool, error)
	Unsubscribe(subscriber domain.UserId, receiver domain.UserId, ctx context.Context) error
	GetSubscriptions(subscriber domain.UserId, ctx context.Context) ([]domain.UserId, error)
	GetSubscribers(user domain.UserId, ctx context.Context) ([]domain.UserId, error)
}

func (p PgDbAdapter) Subscribe(subscriber domain.UserId, receiver domain.UserId, ctx context.Context) (subscribed bool, err error) {
	t, err := p.dbPool.Begin(ctx)
	if err != nil {
		t.Rollback(ctx)
		return false, err
	}

	c, err := t.Exec(ctx, "INSERT INTO subscriptions (follower_id, followed_id)  VALUES ($1, $2) ON CONFLICT DO NOTHING", subscriber, receiver)
	if err != nil {
		t.Rollback(ctx)
		return false, err
	}

	return c.RowsAffected() > 0, t.Commit(ctx)
}

func (p PgDbAdapter) Unsubscribe(subscriber domain.UserId, receiver domain.UserId, ctx context.Context) error {
	t, err := p.dbPool.Begin(ctx)
	if err != nil {
		t.Rollback(ctx)
		return err
	}

	_, err = t.Exec(ctx, "DELETE FROM subscriptions WHERE follower_id=$1 AND followed_id=$2", subscriber, receiver)
	if err != nil {
		t.Rollback(ctx)
		return err
	}

	return t.Commit(ctx)
}

func (p PgDbAdapter) GetSubscriptions(subscriber domain.UserId, ctx context.Context) (subscriptions []domain.UserId, err error) {
	t, err := p.dbPool.Begin(ctx)
	if err != nil {
		t.Rollback(ctx)
		return subscriptions, err
	}

	err = t.QueryRow(ctx, "SELECT followed_id FROM subscriptions WHERE follower_id = $1", subscriber).Scan(&subscriptions)

	return subscriptions, t.Commit(ctx)
}

func (p PgDbAdapter) GetSubscribers(user domain.UserId, ctx context.Context) (subscribers []domain.UserId, err error) {
	t, err := p.dbPool.Begin(ctx)
	if err != nil {
		t.Rollback(ctx)
		return subscribers, err
	}

	err = t.QueryRow(ctx, "SELECT follower_id FROM subscriptions WHERE followed_id = $1", user).Scan(&subscribers)

	return subscribers, t.Commit(ctx)
}
