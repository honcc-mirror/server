package locations

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

type LocationEnum string

const (
	Database LocationEnum = "database"
)

const appRootDirName = "honcc"

var relativeLocations = map[LocationEnum]string{
	Database: "data/honcc_0.0.1.db",
}

func Get(location LocationEnum) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Could not fetch user home dir...")
		return homeDir, nil
	}

	relLocation := relativeLocations[location]
	if relLocation == "" {
		return relLocation, errors.New("unknown location to get")
	}

	locationPath := filepath.Join(homeDir, appRootDirName, relLocation)
	return locationPath, nil
}
