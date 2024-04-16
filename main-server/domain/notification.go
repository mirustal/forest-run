package domain

import "time"

type (
	NotificationType   int
	NotificationStatus int
)

type Notification struct {
	FromUser  UserId
	ToUser    UserId
	Type      NotificationType
	CreatedAt time.Time
	Status    NotificationStatus
}

const (
	NewSubscriberNotification NotificationType = 1
)
