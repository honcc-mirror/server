package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

const (
	testDbFileName = "usersdata_test.db"
)

func cleanupDb() {
	err := os.Remove(testDbFileName)
	if err != nil {
		log.Println("Could not delete test db file!", err)
	}
}

func TestMigration(t *testing.T) {
	t.Cleanup(cleanupDb)
	db, err := sql.Open("sqlite3", testDbFileName)
	if err != nil {
		t.Fatalf("Could not create test database: %s Received error: %s", testDbFileName, err)
	}

	usersData := UsersData{db: db}
	migrateErr := usersData.Migrate()
	if migrateErr != nil {
		t.Fatalf("Could not migrate test database. Received error: %s", migrateErr)
	}
}
