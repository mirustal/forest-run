package boot

import (
	"go.uber.org/zap"
	"main/database"
)

func NewDB(env *Env, logger *zap.Logger) (database.DbAdapter, error) {
	return database.NewAdapter(env, logger)
}
