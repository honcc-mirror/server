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

	for locationName, relLocationPath := range relativeLocations {
		expectedLocation := filepath.Join(homeDir, appRootDirName, relLocationPath)
		location := Get(locationName)

		if location != expectedLocation {
			t.Fatalf(`Received unexpected location.
				Received: %s
				Expected: %s`, location, expectedLocation)
		}
	}
}
