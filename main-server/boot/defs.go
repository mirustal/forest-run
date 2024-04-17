package boot

import (
	"encoding/json"
	"main-server/defs"
	"os"
)

func LoadDefs(env Env) (defs.Defs, error) {
	bytes, err := os.ReadFile(env.DefsPath)
	if err != nil {
		return defs.Defs{}, err
	}

	var d defs.Defs
	err = json.Unmarshal(bytes, &d)

	return d, err
}
