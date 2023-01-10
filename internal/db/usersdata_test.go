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

func setupTestDb(test *testing.T) *sql.DB {
	test.Cleanup(cleanupTestDb)

	db, err := sql.Open("sqlite3", testDbFileName)
	if err != nil {
		test.Fatalf("Could not create test database: %s Received error: %s", testDbFileName, err)
	}

	return db
}

func cleanupTestDb() {
	err := os.Remove(testDbFileName)
	if err != nil {
		log.Println("Could not delete test db file!", err)
	}
}

func TestMigration(test *testing.T) {
	db := setupTestDb(test)
	usersData := UsersData{db: db}
	migrateErr := usersData.Migrate()
	if migrateErr != nil {
		test.Fatalf("Could not migrate test database. Received error: %s", migrateErr)
	}
}
