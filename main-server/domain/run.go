package domain

import "time"

type (
	RunParticipationFormat int
	RunStatus              int
)

type (
	RunId int
	Run   struct {
		Id                    RunId                  `json:"id,omitempty"`
		Name                  string                 `json:"name,omitempty"`
		Description           string                 `json:"description,omitempty"`
		OfficialSiteUrl       string                 `json:"official_site_url,omitempty"`
		AvatarUrl             string                 `json:"avatar_url,omitempty"`
		Route                 Route                  `json:"route"`
		StartTime             time.Time              `json:"start_time"`
		RegistrationUntil     time.Time              `json:"registration_until"`
		MaxParticipants       int                    `json:"max_participants,omitempty"`
		MaxParticipantsOnline int                    `json:"max_participants_online,omitempty"`
		IsPhotoAllowed        bool                   `json:"is_photo_allowed,omitempty"`
		IsVideoAllowed        bool                   `json:"is_video_allowed,omitempty"`
		Status                RunStatus              `json:"status,omitempty"`
		ParticipationFormat   RunParticipationFormat `json:"participation_format,omitempty"`
	}
	Route struct {
		Points []Point `json:"points,omitempty"`
	}
	Point struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
)
