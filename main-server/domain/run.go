package domain

import "time"

type (
	RunId int
	Run   struct {
		Id                  RunId
		Name                string
		OfficialSiteUrl     string
		AvatarUrl           string
		Route               Route
		StartTime           time.Time
		RegistrationUntil   time.Time
		MaxParticipants     int
		IsPhotoAllowed      bool
		IsVideoAllowed      bool
		Status              int
		ParticipationFormat int
	}
	Route struct {
	}
)
