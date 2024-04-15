package notifications

import (
	"main-server/database"
	"main-server/domain"
)

type Manager interface {
	Send(notification domain.Notification) error
	Consume(consumer domain.UserId) ([]domain.Notification, error)
}

func NewManager(db database.DbAdapter) Manager {
	panic("not implemented")
}
