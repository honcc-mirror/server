package users

import (
	"database/sql"
	"fmt"
	"log"

	"honcc/server/internal/db"
)

const usersTableName = "users"

type usersDataTableColumns struct {
	Id          string
	Displayname string
	Username    string
	Key         string
	Salt        string
}

func newUsersDataTableColumns() usersDataTableColumns {
	return usersDataTableColumns{
		Id:          "id",
		Displayname: "displayname",
		Username:    "username",
		Key:         "key",
		Salt:        "salt",
	}
}

type UsersData struct {
	backend db.DBBackend
}

func (usersData *UsersData) Migrate() error {
	columns := newUsersDataTableColumns()
	query := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s (
            %s  INTEGER     PRIMARY KEY AUTOINCREMENT,
            %s  TEXT,
            %s  TEXT        NOT NULL UNIQUE,
            %s  TEXT        NOT NULL,
            %s  TEXT        NOT NULL
        );
    `, usersTableName,
		columns.Id,
		columns.Displayname,
		columns.Username,
		columns.Key,
		columns.Salt)

	_, err := usersData.backend.DB().Exec(query)
	return err
}

func (usersData *UsersData) Create(user User) (*User, error) {
	query := fmt.Sprintf("INSERT INTO %s VALUES (null, ?, ?, ?, ?)", usersTableName)
	result, err := usersData.backend.DB().Exec(
		query,
		user.DisplayName,
		user.Username,
		user.Key,
		user.Salt,
	)

	if err != nil {
		log.Printf(`Could not create new user %#v.
            Received error: %s`, user, err)
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf(`Could not retrieve id for newly created user %#v.
            Received error: %s`, user, err)
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
	columns := newUsersDataTableColumns()
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ?", usersTableName, columns.Id)
	row := usersData.backend.DB().QueryRow(query, userId)
	return scanRowToUser(row)
}

func (usersData *UsersData) Update(user User) (*User, error) {
	columns := newUsersDataTableColumns()
	query := fmt.Sprintf(`
        UPDATE %s SET %s = ?, %s = ?, %s = ?, %s = ?
        WHERE %s = ?`,
		usersTableName,
		columns.Displayname,
		columns.Username,
		columns.Key,
		columns.Salt,
		columns.Id)
	result, err := usersData.backend.DB().Exec(query,
		user.DisplayName,
		user.Username,
		user.Key,
		user.Salt,
		user.Id)
	if err != nil {
		log.Printf(`Could not update user: %#v
            Received error: %s`, user, err)
		return nil, err
	}

	rowsAffected, rowsAffectedErr := result.RowsAffected()
	if rowsAffectedErr != nil {
		log.Printf(`Could not retrieve rows affected for user update.
            User: %#v
            Received error: %s`, user, rowsAffectedErr)
		return nil, rowsAffectedErr
	}

	if rowsAffected == 0 {
		log.Printf(`No rows were affected in user update.
            User with id likely does not exist.
            User: %#v`, user)
		return nil, sql.ErrNoRows
	}

	return &user, nil
}

func (usersData *UsersData) Delete(userId int64) error {
	columns := newUsersDataTableColumns()
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", usersTableName, columns.Id)
	result, err := usersData.backend.DB().Exec(query, userId)
	if err != nil {
		log.Printf(`Could not delete user with id: %d
            Received error: %s`, userId, err)
		return err
	}

	rowsAffected, rowsAffectedErr := result.RowsAffected()
	if rowsAffectedErr != nil {
		log.Printf(`Could not retrieve rows affected for user delete.
            User id: %d
            Received error: %s`, userId, rowsAffectedErr)
		return rowsAffectedErr
	}

	if rowsAffected == 0 {
		log.Printf(`No rows were affected in user update.
            User with id "%d" likely does not exist.`, userId)
		return sql.ErrNoRows
	}

	return nil
}
