package notifications

import (
	"forest-run/common/runs"
	"forest-run/main-server/domain"
	"time"
)

const (
	NewSubscriber                 domain.NotificationType = 1
	NewRunCreated                 domain.NotificationType = 2
	RunChangedStartTime           domain.NotificationType = 3
	RunChangedStartPlace          domain.NotificationType = 4
	RunChangedStatus              domain.NotificationType = 5
	RunChangedRegistrationEndTime domain.NotificationType = 6
)

type RunCreatedBody struct {
	RunId runs.Id `json:"runId"`
}

type RunChangedStartTimeBody struct {
	RunId runs.Id   `json:"runId"`
	Old   time.Time `json:"old"`
	New   time.Time `json:"new"`
}

type RunChangedStartPlaceBody struct {
	RunId runs.Id    `json:"runId"`
	Old   runs.Place `json:"old"`
	New   runs.Place `json:"new"`
}

type RunChangedStatusBody struct {
	RunId runs.Id    `json:"runId"`
	New   domain.Run `json:"new"`
}

type RunChangedRegistrationEndTimeBody struct {
	RunId runs.Id   `json:"runId"`
	Old   time.Time `json:"old"`
	New   time.Time `json:"new"`
}

type EmptyBody struct {
}
