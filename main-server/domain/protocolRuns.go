package domain

import "time"

type (
	CreateRunRequest struct {
		PermissionsTransactionId *TransactionId         `json:"transaction_id"`
		Name                     string                 `json:"name"`
		Description              *string                `json:"description,omitempty"`
		OfficialSiteUrl          *string                `json:"official_site_url,omitempty"`
		AvatarUrl                *string                `json:"avatar_url,omitempty"`
		Route                    Route                  `json:"route"`
		StartTime                time.Time              `json:"start_time"`
		StartPlace               Place                  `json:"start_place"`
		RegistrationUntil        *time.Time             `json:"registration_until,omitempty"`
		MaxParticipants          int                    `json:"max_participants"`
		RunPermissions           RunPermissionsType     `json:"permissions"`
		ParticipationFormat      RunParticipationFormat `json:"participation_format,omitempty"`
	}
	CreateRunResponse struct {
		Run Run `json:"run"`
	}
)
