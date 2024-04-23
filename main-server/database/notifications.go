package database

import (
	"context"
	"forest-run/common"
	"forest-run/main-server/domain"
	"github.com/jackc/pgx/v5"
)

type NotificationsRepo interface {
	Store(notification domain.Notification, ctx context.Context) error
	GetNotifications(userId common.UserId, ctx context.Context) ([]domain.Notification, error)
	StoreMany(notifications []domain.Notification, ctx context.Context) error
}

func (p PgDbAdapter) Store(notification domain.Notification, ctx context.Context) error {
	t, err := p.dbPool.Begin(ctx)
	if err != nil {
		t.Rollback(ctx)
		return err
	}

	_, err = t.Exec(ctx, "INSERT INTO notifications (from_user_id, to_user_id, type, body) VALUES ($1, $2, $3, $4)", notification.FromUser, notification.ToUser, notification.Type, notification.Body)
	if err != nil {
		t.Rollback(ctx)
		return err
	}

	return t.Commit(ctx)
}

func (p PgDbAdapter) GetNotifications(userId common.UserId, ctx context.Context) ([]domain.Notification, error) {
	t, err := p.dbPool.Begin(ctx)
	if err != nil {
		t.Rollback(ctx)
		return nil, err
	}

	var notifications []domain.Notification
	err = t.QueryRow(ctx, "SELECT * FROM notifications WHERE to_user_id = $1", userId).Scan(&notifications)
	if err != nil {
		t.Rollback(ctx)
		return nil, err
	}

	return notifications, t.Commit(ctx)
}

func (p PgDbAdapter) StoreMany(notifications []domain.Notification, ctx context.Context) error {
	t, err := p.dbPool.Begin(ctx)
	if err != nil {
		t.Rollback(ctx)
		return err
	}

	_, err = t.CopyFrom(ctx, pgx.Identifier{"notifications"}, []string{"from_user_id", "to_user_id", "type", "body"}, pgx.CopyFromSlice(len(notifications), func(i int) ([]any, error) {
		return []any{notifications[i].FromUser, notifications[i].ToUser, notifications[i].Type, notifications[i].Body}, nil
	}))

	if err != nil {
		t.Rollback(ctx)
		return err
	}

	return t.Commit(ctx)
}
