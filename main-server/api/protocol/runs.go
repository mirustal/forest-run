package protocol

import (
	"forest-run/common"
	"forest-run/common/runs"
	"forest-run/main-server/domain"
	"time"
)

type (
	CreateRunRequest struct {
		PermissionsTransactionId *domain.TransactionId         `json:"transaction_id" validate:"optional"`
		Name                     string                        `json:"name"`
		Description              *string                       `json:"description,omitempty" validate:"optional"`
		OfficialSiteUrl          *string                       `json:"official_site_url,omitempty" validate:"optional"`
		AvatarUrl                *string                       `json:"avatar_url,omitempty" validate:"optional"`
		Route                    runs.Route                    `json:"route"`
		StartTime                time.Time                     `json:"start_time"`
		StartPlace               runs.Place                    `json:"start_place"`
		RegistrationUntil        *time.Time                    `json:"registration_until,omitempty" validate:"optional"`
		MaxParticipants          int                           `json:"max_participants"`
		RunPermissions           runs.PermissionsType          `json:"permissions"`
		ParticipationFormat      domain.RunParticipationFormat `json:"participation_format,omitempty"`
	}
	CreateRunResponse struct {
		Run domain.Run `json:"run"`
	}
)

type (
	UpdateRunRequest struct {
		RunId                    runs.Id                        `json:"run_id"`
		PermissionsTransactionId *domain.TransactionId          `json:"transaction_id" validate:"optional"`
		Name                     *string                        `json:"name" validate:"optional"`
		Description              *string                        `json:"description,omitempty" validate:"optional"`
		OfficialSiteUrl          *string                        `json:"official_site_url,omitempty" validate:"optional"`
		AvatarUrl                *string                        `json:"avatar_url,omitempty" validate:"optional"`
		Route                    *runs.Route                    `json:"route" validate:"optional"`
		StartTime                *time.Time                     `json:"start_time" validate:"optional"`
		StartPlace               *runs.Place                    `json:"start_place" validate:"optional"`
		RegistrationUntil        *time.Time                     `json:"registration_until,omitempty" validate:"optional"`
		MaxParticipants          *int                           `json:"max_participants" validate:"optional"`
		RunPermissions           *runs.PermissionsType          `json:"permissions" validate:"optional"`
		ParticipationFormat      *domain.RunParticipationFormat `json:"participation_format,omitempty" validate:"optional"`
	}
	UpdateRunResponse struct {
		Run domain.Run `json:"run"`
	}
)

type (
	InviteRunRequest struct {
		RunId runs.Id         `json:"run_id"`
		Users []common.UserId `json:"users"`
	}
	InviteRunResponse struct {
	}
)
