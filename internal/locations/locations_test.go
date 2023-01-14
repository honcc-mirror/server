package locations

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGet(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("Could not fetch user home dir, received error: %s", err)
	}

	expectedLocation := filepath.Join(homeDir, appRootDirName, relativeLocations[Database])
	location, err := Get(Database)
	if err != nil {
		t.Fatalf("Could not acquire location, received error: %s", err)
	}

	if location != expectedLocation {
		t.Fatalf(`Received unexpected location.
			Received: %s
			Expected: %s`, location, expectedLocation)
	}
}
