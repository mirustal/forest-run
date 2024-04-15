package domain

import "time"

type Notification struct {
	FromUser  UserId
	ToUser    UserId
	Type      int
	CreatedAt time.Time
	Status    int
}

const (
	NewSubscriberNotification = 1
)
