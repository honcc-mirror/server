package db

import (
	"database/sql"
)

type DBBackend interface {
	DB() *sql.DB
	Open(string) error
}
