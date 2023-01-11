package db

import (
	"database/sql"
	"fmt"
	"log"
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

func (usersData *UsersData) Create(user User) (*User, error) {
	query := fmt.Sprintf("INSERT INTO %s VALUES (null, ?, ?, ?, ?)", usersTableName)
	result, err := usersData.db.Exec(
		query,
		user.DisplayName,
		user.Username,
		user.Key,
		user.Salt,
	)

	if err != nil {
		log.Printf("Could not create new user %#v, received error: %s", user, err)
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Could not retrieve id for newly created user %#v, received error: %s", user, err)
		return nil, err
	}

	user.Id = id
	return &user, err
}

func scanRowToUser(row *sql.Row) (*User, error) {
	var user User
	err := row.Scan(&user.Id, &user.DisplayName, &user.Username, &user.Key, &user.Salt)
	if err != nil {
		log.Printf("Could not retrieve user: %s", err)
		return nil, err
	}

	return &user, err
}

func (usersData *UsersData) UserFromId(userId int64) (*User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ?", usersTableName, idColumn)
	row := usersData.db.QueryRow(query, userId)
	return scanRowToUser(row)
}
