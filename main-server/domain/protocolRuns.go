package domain

import "time"

type (
	CreateRunRequest struct {
		PermissionsTransactionId *TransactionId         `json:"transaction_id" validate:"optional"`
		Name                     string                 `json:"name"`
		Description              *string                `json:"description,omitempty" validate:"optional"`
		OfficialSiteUrl          *string                `json:"official_site_url,omitempty" validate:"optional"`
		AvatarUrl                *string                `json:"avatar_url,omitempty" validate:"optional"`
		Route                    Route                  `json:"route"`
		StartTime                time.Time              `json:"start_time"`
		StartPlace               Place                  `json:"start_place"`
		RegistrationUntil        *time.Time             `json:"registration_until,omitempty" validate:"optional"`
		MaxParticipants          int                    `json:"max_participants"`
		RunPermissions           RunPermissionsType     `json:"permissions"`
		ParticipationFormat      RunParticipationFormat `json:"participation_format,omitempty"`
	}
	CreateRunResponse struct {
		Run Run `json:"run"`
	}
)

type (
	UpdateRunRequest struct {
		RunId                    RunId                   `json:"run_id"`
		PermissionsTransactionId *TransactionId          `json:"transaction_id" validate:"optional"`
		Name                     *string                 `json:"name" validate:"optional"`
		Description              *string                 `json:"description,omitempty" validate:"optional"`
		OfficialSiteUrl          *string                 `json:"official_site_url,omitempty" validate:"optional"`
		AvatarUrl                *string                 `json:"avatar_url,omitempty" validate:"optional"`
		Route                    *Route                  `json:"route" validate:"optional"`
		StartTime                *time.Time              `json:"start_time" validate:"optional"`
		StartPlace               *Place                  `json:"start_place" validate:"optional"`
		RegistrationUntil        *time.Time              `json:"registration_until,omitempty" validate:"optional"`
		MaxParticipants          *int                    `json:"max_participants" validate:"optional"`
		RunPermissions           *RunPermissionsType     `json:"permissions" validate:"optional"`
		ParticipationFormat      *RunParticipationFormat `json:"participation_format,omitempty" validate:"optional"`
	}
	UpdateRunResponse struct {
		Run Run `json:"run"`
	}
)
