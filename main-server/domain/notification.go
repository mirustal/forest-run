package domain

type Notification struct {
	FromUser UserId
	ToUser   UserId
	Type     int
}

const (
	NewSubscriberNotification = 1
)
