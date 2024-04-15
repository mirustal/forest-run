package notifications

import (
	"main-server/database"
	"main-server/domain"
)

type Manager interface {
	Send(notification domain.Notification) error
	Consume(consumer domain.UserId) ([]domain.Notification, error)
}

type manager struct {
	db database.DbAdapter
}

func NewManager(db database.DbAdapter) Manager {
	return &manager{db: db}
}

func (m manager) Send(notification domain.Notification) error {
	//TODO implement me
	return nil
}

func (m manager) Consume(consumer domain.UserId) ([]domain.Notification, error) {
	//TODO implement me
	return []domain.Notification{}, nil
}
