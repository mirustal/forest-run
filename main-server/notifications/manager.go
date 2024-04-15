package notifications

import (
	"context"
	"main-server/database"
	"main-server/domain"
)

type Manager interface {
	Send(notification domain.Notification, ctx context.Context) error
	Consume(consumer domain.UserId, ctx context.Context) ([]domain.Notification, error)
}

type manager struct {
	db database.DbAdapter
}

func NewManager(db database.DbAdapter) Manager {
	return &manager{db: db}
}

func (m manager) Send(notification domain.Notification, ctx context.Context) error {
	return m.db.Store(notification, ctx)
}

func (m manager) Consume(consumer domain.UserId, ctx context.Context) ([]domain.Notification, error) {
	return m.db.GetNotifications(consumer, ctx)
}
