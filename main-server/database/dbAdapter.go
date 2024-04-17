package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pashagolub/pgxmock/v3"
	"go.uber.org/zap"
	"main-server/boot"
)

type DbAdapter interface {
	AuthRepo
	SubscriptionsRepo
	NotificationsRepo
	RunsRepo
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
	dbPool PgxPool
	logger *zap.Logger
}

type PgxPool interface {
	Begin(context.Context) (pgx.Tx, error)
	Close()
}

type MockDbAdapter struct {
	PgDbAdapter
	Mock pgxmock.PgxPoolIface
}

func NewMockAdapter() (MockDbAdapter, error) {
	dbPool, err := pgxmock.NewPool()
	if err != nil {
		return MockDbAdapter{}, err
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		return MockDbAdapter{}, err
	}

	return MockDbAdapter{
		PgDbAdapter: PgDbAdapter{
			dbPool: dbPool,
			logger: logger,
		},
		mock: dbPool,
	}, nil
}
