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

var (
	testUser = User{
		Id:          -1,
		DisplayName: "Test user",
		Username:    "test",
		Key:         "ABC",
		Salt:        "CDE",
	}
)

// Helper methods for setting up and breaking down testing env

func cleanupTestDb() {
	err := os.Remove(testDbFileName)
	if err != nil {
		log.Println("Could not delete test db file!", err)
	}
}

// Tests

func TestMigration(test *testing.T) {
	test.Cleanup(cleanupTestDb)

	db, err := sql.Open("sqlite3", testDbFileName)
	if err != nil {
		test.Fatalf(`Could not create test database: %s
			Received error: %s`, testDbFileName, err)
	}

	usersData := UsersData{db: db}
	migrateErr := usersData.Migrate()
	if migrateErr != nil {
		test.Fatalf(`Could not migrate test database.
			Received error: %s`, migrateErr)
	}
}

func TestCreate(test *testing.T) {
	test.Cleanup(cleanupTestDb)

	db, err := sql.Open("sqlite3", testDbFileName)
	if err != nil {
		test.Fatalf(`Could not create test database: %s
			Received error: %s`, testDbFileName, err)
	}

	usersData := UsersData{db: db}
	migrateErr := usersData.Migrate()
	if migrateErr != nil {
		test.Fatalf("Could not migrate test database")
	}

	createdUser, createErr := usersData.Create(testUser)
	if createErr != nil {
		test.Fatalf("Could not create user, received error %s", createErr)
	}

	if createdUser.Id == testUser.Id {
		test.Fatalf(`Newly created user has invalid id value.
			New user should have a value different to the input test user.
			Test user id: %d
			Created user id: %d`, testUser.Id, createdUser.Id)
	}
}

func TestUserFromId(test *testing.T) {
	test.Cleanup(cleanupTestDb)

	db, err := sql.Open("sqlite3", testDbFileName)
	if err != nil {
		test.Fatalf(`Could not create test database: %s
			Received error: %s`, testDbFileName, err)
	}

	usersData := UsersData{db: db}
	migrateErr := usersData.Migrate()
	if migrateErr != nil {
		test.Fatalf("Could not migrate test database")
	}

	createdUser, createErr := usersData.Create(testUser)
	if createErr != nil {
		test.Fatalf("Could not create user, received error %s", createErr)
	}

	returnedUser, returnedErr := usersData.UserFromId(createdUser.Id)
	if returnedErr != nil {
		test.Fatalf("Could not read user with id: %s", returnedErr)
	}

	if *returnedUser != *createdUser {
		test.Fatalf(`Received user data differs from expected input user data.
			In user data: %#v
			Out user data: %#v`, createdUser, returnedUser)
	}
}

func TestUpdate(test *testing.T) {
	test.Cleanup(cleanupTestDb)

	db, err := sql.Open("sqlite3", testDbFileName)
	if err != nil {
		test.Fatalf(`Could not create test database: %s
			Received error: %s`, testDbFileName, err)
	}

	usersData := UsersData{db: db}
	migrateErr := usersData.Migrate()
	if migrateErr != nil {
		test.Fatalf("Could not migrate test database")
	}

	createdUser, createErr := usersData.Create(testUser)
	if createErr != nil {
		test.Fatalf("Could not create user, received error %s", createErr)
	}

	modifiedUser := User{
		Id:          createdUser.Id,
		DisplayName: "Modified user",
		Username:    "modified_user",
		Key:         "JKL",
		Salt:        "XYZ",
	}

	updatedUser, updateErr := usersData.Update(modifiedUser)
	if updateErr != nil {
		test.Fatalf("Could not update user, received error %s", updateErr)
	}

	if *updatedUser != modifiedUser {
		test.Fatalf(`Updated user data differs from expected.
			Initial user data: %#v
			Expected user data: %#v
			Out user data: %#v`, createdUser, modifiedUser, updatedUser)
	}
}
