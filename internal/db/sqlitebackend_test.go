package db

import (
	"log"
	"os"
	"testing"
)

const testDbFileName = "testdb.db"

func cleanupTestDb() {
	err := os.Remove(testDbFileName)
	if err != nil {
		log.Println("Could not delete test db file!", err)
	}
}

func TestSQLiteBackendOpen(t *testing.T) {
	t.Cleanup(cleanupTestDb)

	backend := SQLiteBackend{}
	err := backend.Open(testDbFileName)
	if err != nil {
		t.Fatalf("Could not open sqlite db, received error: %s", err)
	}

	pingErr := backend.db.Ping()
	if pingErr != nil {
		t.Fatalf("Received error on ping attempt: %s", pingErr)
	}
}
