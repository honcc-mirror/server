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

func TestCreate(test *testing.T) {
	db := setupTestDb(test)
	usersData := UsersData{db: db}
	migrateErr := usersData.Migrate()
	if migrateErr != nil {
		test.Fatalf("Could not migrate test database")
	}

	testUser := User{
		Id:          -1,
		DisplayName: "Test user",
		Username:    "test",
		Key:         "ABC",
		Salt:        "CDE",
	}

	createdUser, createErr := usersData.Create(testUser)
	if createErr != nil {
		test.Fatalf("Could not create user, received error %s", createErr)
	}

	if createdUser.Id == testUser.Id {
		test.Fatalf("Newly created user should not have the invalid value of the test user. Test user id: %d, Created user id: %d", testUser.Id, createdUser.Id)
	}
}
