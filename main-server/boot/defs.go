package boot

import (
	"main-server/defs"
	"main-server/domain"
)

func LoadDefs(env Env) (defs.Defs, error) {
	// todo: move to json file, better to cdn
	return defs.Defs{
		RunPermissionsDefs: defs.RunPermissionsDefs{
			Types: map[domain.RunPermissionsType]defs.RunPermissionsDef{
				domain.RunPermissionsType(0): {
					MaxOnlineParticipants: 10,
					IsPhotoAllowed:        false,
					IsStoriesAllowed:      false,
					IsStreamingAllowed:    false,
				},
				domain.RunPermissionsType(10): {
					MaxOnlineParticipants: -1,
					IsPhotoAllowed:        true,
					IsStoriesAllowed:      true,
					IsStreamingAllowed:    true,
				},
			},
		},
	}, nil
}
