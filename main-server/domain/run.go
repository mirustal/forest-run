package domain

import (
	"forest-run/common"
	"forest-run/common/runs"
	"time"
)

type (
	RunParticipationFormat int
	RunStatus              int
)

const (
	OpenRunFormat   = RunParticipationFormat(0)
	ClosedRunFormat = RunParticipationFormat(1)
)

const (
	OpenRunStatus     = RunStatus(0)
	RunningRunStatus  = RunStatus(1)
	FinishedRunStatus = RunStatus(2)
	CanceledRunStatus = RunStatus(4)
)

type Run struct {
	runs.EssentialInfo
	Creator             common.UserId          `json:"creator,omitempty"`
	Name                string                 `json:"name,omitempty"`
	Description         *string                `json:"description,omitempty"`
	OfficialSiteUrl     *string                `json:"official_site_url,omitempty"`
	AvatarUrl           *string                `json:"avatar_url,omitempty"`
	RegistrationUntil   time.Time              `json:"registration_until"`
	MaxParticipants     int                    `json:"max_participants,omitempty"`
	Status              RunStatus              `json:"status,omitempty"`
	ParticipationFormat RunParticipationFormat `json:"participation_format,omitempty"`
}
