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
	defer func() {
		if err != nil {
			_ = t.Rollback(ctx)
		}
	}()

	if err != nil {
		return subscriptions, err
	}

	r, err := t.Query(ctx, "SELECT followed_id FROM subscriptions WHERE follower_id = $1", subscriber)

	if err != nil {
		return subscriptions, err
	}

	subscriptions = make([]domain.UserId, 0)
	for r.Next() {
		if err = r.Err(); err != nil {
			return subscriptions, err
		}
		var sub domain.UserId
		err = r.Scan(&sub)
		if err != nil {
			return subscriptions, err
		}
		subscriptions = append(subscriptions, sub)
	}

	if err != nil {
		return subscriptions, err
	}

	return subscriptions, t.Commit(ctx)
}

func (p PgDbAdapter) GetSubscribers(user domain.UserId, ctx context.Context) (subscribers []domain.UserId, err error) {
	t, err := p.dbPool.Begin(ctx)
	defer func() {
		if err != nil {
			_ = t.Rollback(ctx)
		}
	}()

	if err != nil {
		return subscribers, err
	}

	r, err := t.Query(ctx, "SELECT follower_id FROM subscriptions WHERE followed_id = $1", user)

	if err != nil {
		return subscribers, err
	}

	subscribers = make([]domain.UserId, 0)
	for r.Next() {
		if err = r.Err(); err != nil {
			return subscribers, err
		}
		var sub domain.UserId
		err = r.Scan(&sub)
		if err != nil {
			return subscribers, err
		}
		subscribers = append(subscribers, sub)
	}

	if err != nil {
		return subscribers, err
	}

	return subscribers, t.Commit(ctx)
}
