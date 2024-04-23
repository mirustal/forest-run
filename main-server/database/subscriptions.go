package database

import (
	"context"
	"forest-run/common"
)

type SubscriptionsRepo interface {
	Subscribe(subscriber common.UserId, receiver common.UserId, ctx context.Context) (bool, error)
	Unsubscribe(subscriber common.UserId, receiver common.UserId, ctx context.Context) error
	GetSubscriptions(subscriber common.UserId, ctx context.Context) ([]common.UserId, error)
	GetSubscribers(user common.UserId, ctx context.Context) ([]common.UserId, error)
}

func (p PgDbAdapter) Subscribe(subscriber common.UserId, receiver common.UserId, ctx context.Context) (subscribed bool, err error) {
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

func (p PgDbAdapter) Unsubscribe(subscriber common.UserId, receiver common.UserId, ctx context.Context) error {
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

func (p PgDbAdapter) GetSubscriptions(subscriber common.UserId, ctx context.Context) (subscriptions []common.UserId, err error) {
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

	subscriptions = make([]common.UserId, 0)
	for r.Next() {
		if err = r.Err(); err != nil {
			return subscriptions, err
		}
		var sub common.UserId
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

func (p PgDbAdapter) GetSubscribers(user common.UserId, ctx context.Context) (subscribers []common.UserId, err error) {
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

	subscribers = make([]common.UserId, 0)
	for r.Next() {
		if err = r.Err(); err != nil {
			return subscribers, err
		}
		var sub common.UserId
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
