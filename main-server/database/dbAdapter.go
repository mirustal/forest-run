package database

import "github.com/jackc/pgx/v5"

type DbAdapter interface {
}

type PgDbAdapter struct {
	conn *pgx.Conn
}

func NewAdapter(conn *pgx.Conn) (DbAdapter, error) {
	dbAdapter := &PgDbAdapter{conn: conn}
	return dbAdapter, nil
}
