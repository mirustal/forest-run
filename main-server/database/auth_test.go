package database

import (
	"context"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"main-server/domain"
	"testing"
)

type MockDbAdapter struct {
	PgDbAdapter
	mock pgxmock.PgxPoolIface
}

func newMockAdapter() (MockDbAdapter, error) {
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

func TestStoreUser(t *testing.T) {
	db, err := newMockAdapter()
	assert.Nil(t, err)

	db.mock.ExpectBegin()
	db.mock.ExpectExec("INSERT INTO users").WithArgs(domain.Username("test"), domain.HashedPassword("test")).WillReturnResult(pgxmock.NewResult("INSERT", 1))
	db.mock.ExpectCommit()

	err = db.StoreNewUser(domain.Username("test"), domain.HashedPassword("test"), context.Background())
	assert.Nil(t, err)

	if db.mock.ExpectationsWereMet() != nil {
		t.Errorf("unfulfilled expectations: %s", db.mock.ExpectationsWereMet())
	}
}
