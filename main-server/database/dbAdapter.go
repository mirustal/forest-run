package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"main-server/boot"
	"main-server/domain"
)

type DbAdapter interface {
	AuthRepo
}

type AuthRepo interface {
	StoreNewUser(username domain.Username, password domain.HashedPassword, ctx context.Context) error
	GetUserByUsername(username domain.Username, ctx context.Context) (domain.User, error)
	StoreUserRefreshToken(id domain.UserId, data domain.RefreshTokenData, ctx context.Context) error
	GetUserRefreshToken(id domain.UserId, ctx context.Context) error
}

type PgDbAdapter struct {
	dbPool *pgxpool.Pool
	logger *zap.Logger
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
