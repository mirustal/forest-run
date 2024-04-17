package domain

import "time"

type (
	RunParticipationFormat int
	RunStatus              int
	RunPermissionsType     int
)

const (
	FreeRunPermissionsType = RunPermissionsType(0)
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

type (
	RunId int
	Run   struct {
		Id                  RunId                  `json:"id,omitempty"`
		Creator             UserId                 `json:"creator,omitempty"`
		Name                string                 `json:"name,omitempty"`
		Description         *string                `json:"description,omitempty"`
		OfficialSiteUrl     *string                `json:"official_site_url,omitempty"`
		AvatarUrl           *string                `json:"avatar_url,omitempty"`
		Route               Route                  `json:"route"`
		StartTime           time.Time              `json:"start_time"`
		StartPlace          Place                  `json:"start_place"`
		RegistrationUntil   time.Time              `json:"registration_until"`
		MaxParticipants     int                    `json:"max_participants,omitempty"`
		RunPermissions      RunPermissionsType     `json:"permissions,omitempty"`
		Status              RunStatus              `json:"status,omitempty"`
		ParticipationFormat RunParticipationFormat `json:"participation_format,omitempty"`
	}
	Place struct {
		Address string `json:"address"`
		Point   Point  `json:"point"`
	}
	Route struct {
		Points []Point `json:"points,omitempty"`
	}
	Point struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
)
