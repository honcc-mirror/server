package db

import (
	"database/sql"
	"honcc/server/internal/locations"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteBackend struct {
	db *sql.DB
}

func (sqlb *SQLiteBackend) DB() *sql.DB {
	return sqlb.db
}

func (sqlb *SQLiteBackend) Open() error {
	db, err := sql.Open("sqlite3", locations.Get(locations.Database))
	if err != nil {
		log.Fatalf("Could not open sqlite database: %s", err)
		return err
	}

	sqlb.db = db

	return err
}
