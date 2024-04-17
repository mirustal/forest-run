package defs

import "main-server/domain"

type RunPermissionsDefs struct {
	Types map[domain.RunPermissionsType]RunPermissionsDef `json:"types,omitempty"`
}

type RunPermissionsDef struct {
	MaxOnlineParticipants int  `json:"max_online_participants,omitempty"`
	IsPhotoAllowed        bool `json:"is_photo_allowed,omitempty"`
	IsStoriesAllowed      bool `json:"is_stories_allowed,omitempty"`
	IsStreamingAllowed    bool `json:"is_streaming_allowed,omitempty"`
}
