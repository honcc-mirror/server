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

func setupUsersData(test *testing.T) *UsersData {
	test.Cleanup(cleanupTestDb)

	db, err := sql.Open("sqlite3", testDbFileName)
	if err != nil {
		test.Fatalf(`Could not create test database: %s
			Received error: %s`, testDbFileName, err)
		return nil
	}

	usersData := UsersData{db: db}
	migrateErr := usersData.Migrate()
	if migrateErr != nil {
		test.Fatalf(`Could not migrate test database.
			Received error: %s`, migrateErr)
		return nil
	}

	return &usersData
}

func createTestUser(test *testing.T, usersData *UsersData) *User {
	createdUser, createErr := usersData.Create(testUser)
	if createErr != nil {
		test.Fatalf("Could not create user, received error %s", createErr)
		return nil
	}

	if createdUser.Id == testUser.Id {
		test.Fatalf(`Newly created user has invalid id value.
			New user should have a value different to the input test user.
			Test user id: %d
			Created user id: %d`, testUser.Id, createdUser.Id)
		return nil
	}

	return createdUser
}

// Tests

func TestMigration(test *testing.T) {
	setupUsersData(test)
}

func TestCreate(test *testing.T) {
	usersData := setupUsersData(test)
	createTestUser(test, usersData)
}

func TestUserFromId(test *testing.T) {
	usersData := setupUsersData(test)
	createdUser := createTestUser(test, usersData)
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

func TestUserFromIdFail(test *testing.T) {
	usersData := setupUsersData(test)

	// Should get error from trying to fetch user in empty db
	_, emptyDbErr := usersData.UserFromId(0)
	if emptyDbErr == nil {
		test.Fatalf("Should receive error after trying to retrieve user from empty db")
	}

	// Should get error from trying to fetch user with wrong id
	createdUser := createTestUser(test, usersData)
	_, wrongIdErr := usersData.UserFromId(createdUser.Id + 1)
	if wrongIdErr == nil {
		test.Fatalf("Should receive error after trying to retrieve user with bad id")
	}
}

func TestUpdate(test *testing.T) {
	usersData := setupUsersData(test)
	createdUser := createTestUser(test, usersData)
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

func TestUpdateFail(test *testing.T) {
	usersData := setupUsersData(test)
	createdUser := createTestUser(test, usersData)

	// Should not be able to update user with an invalid id
	invalidModifiedUser := User{
		Id:          createdUser.Id + 1,
		DisplayName: "Modified user",
		Username:    "modified_user",
		Key:         "JKL",
		Salt:        "XYZ",
	}

	_, invalidUpdateErr := usersData.Update(invalidModifiedUser)
	if invalidUpdateErr == nil {
		test.Fatalf("Should not be able to update user with an invalid id")
	}
}

func TestDelete(test *testing.T) {
	usersData := setupUsersData(test)
	createdUser := createTestUser(test, usersData)
	deleteErr := usersData.Delete(createdUser.Id)
	if deleteErr != nil {
		test.Fatalf("Could not delete user, received error %s", deleteErr)
	}
}

func TestDeleteFail(test *testing.T) {
	usersData := setupUsersData(test)
	createdUser := createTestUser(test, usersData)

	// Should not be able to delete a user using an invalid id
	invalidIdErr := usersData.Delete(createdUser.Id + 1)
	if invalidIdErr == nil {
		test.Fatalf("Should not be able to delete user with invalid id")
	}
}
