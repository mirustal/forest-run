package domain

import (
	"encoding/json"
	"forest-run/common"
	"time"
)

type (
	NotificationType   int
	NotificationStatus int
)

type Notification struct {
	FromUser  common.UserId
	ToUser    common.UserId
	Type      NotificationType
	CreatedAt time.Time
	Status    NotificationStatus
	Body      string
}

func (n Notification) WithBody(b any) (Notification, error) {
	body, err := json.Marshal(b)
	if err != nil {
		return n, err
	}

	n.Body = string(body)
	return n, nil
}
