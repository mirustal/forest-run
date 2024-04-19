package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"main-server/domain"
)

type NotificationsRepo interface {
	Store(notification domain.Notification, ctx context.Context) error
	GetNotifications(userId domain.UserId, ctx context.Context) ([]domain.Notification, error)
	StoreMany(notifications []domain.Notification, ctx context.Context) error
}

func (p PgDbAdapter) Store(notification domain.Notification, ctx context.Context) error {
	t, err := p.dbPool.Begin(ctx)
	if err != nil {
		t.Rollback(ctx)
		return err
	}

	_, err = t.Exec(ctx, "INSERT INTO notifications (from_user_id, to_user_id, type) VALUES ($1, $2, $3)", notification.FromUser, notification.ToUser, notification.Type)
	if err != nil {
		t.Rollback(ctx)
		return err
	}

	return t.Commit(ctx)
}

func (p PgDbAdapter) GetNotifications(userId domain.UserId, ctx context.Context) ([]domain.Notification, error) {
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

	_, err = t.CopyFrom(ctx, pgx.Identifier{"notifications"}, []string{"from_user_id", "to_user_id", "type"}, pgx.CopyFromSlice(len(notifications), func(i int) ([]any, error) {
		return []any{notifications[i].FromUser, notifications[i].ToUser, notifications[i].Type}, nil
	}))

	if err != nil {
		t.Rollback(ctx)
		return err
	}

	return t.Commit(ctx)
}
