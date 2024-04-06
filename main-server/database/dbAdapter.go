package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"main-server/boot"
	"main-server/domain"
)

type DbAdapter interface {
	DropAuthTokens(ctx context.Context, id domain.UserId) error
	GetLatestToken(ctx context.Context, id domain.UserId) (*domain.AuthTokenData, error)
	StoreToken(ctx context.Context, id domain.UserId, token domain.AuthToken) error
}

type PgDbAdapter struct {
	dbPool *pgxpool.Pool
	logger *zap.Logger
}

func NewAdapter(env *boot.Env, logger *zap.Logger) (DbAdapter, error) {
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
