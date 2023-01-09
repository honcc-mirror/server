package db

import (
	"database/sql"
)

type UsersData struct {
	db *sql.DB
}

func (usersData *UsersData) Migrate() error {
	query := `
        CREATE TABLE IF NOT EXISTS users (
            id          INTEGER     PRIMARY KEY AUTOINCREMENT,
            displayname TEXT,
            username    TEXT        NOT NULL UNIQUE,
            key         TEXT        NOT NULL,
            salt        TEXT        NOT NULL
        );
    `

	_, err := usersData.db.Exec(query)
	return err
}
