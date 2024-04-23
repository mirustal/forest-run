package defs

import (
	"encoding/json"
	"forest-run/common/configs"
	"os"
)

func Load(env configs.CommonConfig) (Defs, error) {
	bytes, err := os.ReadFile(env.DefsPath)
	if err != nil {
		return Defs{}, err
	}

	var d Defs
	err = json.Unmarshal(bytes, &d)

	return d, err
}
