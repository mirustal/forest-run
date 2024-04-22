package domain

import (
	"encoding/json"
	"time"
)

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

const (
	NewSubscriberNotification        NotificationType = 1
	NewRunCreatedNotification        NotificationType = 2
	RunChangedStartTimeNotification  NotificationType = 3
	RunChangedStartPlaceNotification NotificationType = 4
	RunChangedStatusNotification     NotificationType = 5
)

type RunCreatedNotificationBody struct {
	RunId RunId `json:"runId"`
}

type RunChangedStartTimeNotificationBody struct {
	RunId RunId     `json:"runId"`
	Old   time.Time `json:"old"`
	New   time.Time `json:"new"`
}

type RunChangedStartPlaceNotificationBody struct {
	RunId RunId `json:"runId"`
	Old   Place `json:"old"`
	New   Place `json:"new"`
}

type RunChangedStatusNotificationBody struct {
	RunId RunId `json:"runId"`
	New   Run   `json:"new"`
}

type EmptyNotificationBody struct {
}
