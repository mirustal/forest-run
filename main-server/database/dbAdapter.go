package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"main-server/boot"
)

type DbAdapter interface {
	AuthRepo
	SubscriptionsRepo
}

func NewAdapter(env boot.DBConfig, logger *zap.Logger) (DbAdapter, error) {
	dbPool, err := pgxpool.New(context.Background(), env.DBUrl)
	if err != nil {
		logger.Fatal("Unable to create connection pool: ", zap.Error(err))
		return nil, err
	}

	return &PgDbAdapter{
		dbPool: dbPool,
		logger: logger,
	}, nil
}

type PgDbAdapter struct {
	dbPool *pgxpool.Pool
	logger *zap.Logger
}
