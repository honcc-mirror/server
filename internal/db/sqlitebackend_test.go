package db

import (
	"testing"
)

func TestSQLiteBackendOpen(t *testing.T) {
	backend := SQLiteBackend{}
	err := backend.Open()
	if err != nil {
		t.Fatalf("Could not open sqlite db, received error: %s", err)
	}

	pingErr := backend.db.Ping()
	if pingErr != nil {
		t.Fatalf("Received error on ping attempt: %s", pingErr)
	}
}
