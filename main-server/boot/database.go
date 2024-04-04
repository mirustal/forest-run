package boot

import (
	"context"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
	"main/database"
)

func NewDB(env *Env, logger *zap.Logger) (database.DbAdapter, error) {
	conn, err := pgx.Connect(context.Background(), env.DBUrl)
	if err != nil {
		logger.Fatal("Can't connect to DB: ", zap.Error(err))
		return nil, err
	}

	return database.NewAdapter(conn)
}
