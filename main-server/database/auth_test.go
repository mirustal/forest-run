package database

import (
	"context"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
	"main-server/domain"
	"testing"
)

func TestStoreUser(t *testing.T) {
	db, err := NewMockAdapter()
	assert.Nil(t, err)

	db.Mock.ExpectBegin()
	db.Mock.ExpectExec("INSERT INTO users").WithArgs(domain.Username("test"), domain.HashedPassword("test")).WillReturnResult(pgxmock.NewResult("INSERT", 1))
	db.Mock.ExpectCommit()

	err = db.StoreNewUser(domain.Username("test"), domain.HashedPassword("test"), context.Background())
	assert.Nil(t, err)

	if err = db.Mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %s", err)
	}
}
