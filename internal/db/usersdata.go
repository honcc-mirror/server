package db

import (
	"database/sql"
	"fmt"
)

const (
	usersTableName    = "users"
	idColumn          = "id"
	displayNameColumn = "displayname"
	userNameColumn    = "username"
	keyColumn         = "key"
	saltColumn        = "salt"
)

type UsersData struct {
	db *sql.DB
}

func (usersData *UsersData) Migrate() error {
	query := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s (
            %s  INTEGER     PRIMARY KEY AUTOINCREMENT,
            %s  TEXT,
            %s  TEXT        NOT NULL UNIQUE,
            %s  TEXT        NOT NULL,
            %s  TEXT        NOT NULL
        );
    `, usersTableName,
		idColumn,
		displayNameColumn,
		userNameColumn,
		keyColumn,
		saltColumn)

	_, err := usersData.db.Exec(query)
	return err
}
