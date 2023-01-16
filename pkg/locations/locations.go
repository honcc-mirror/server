package locations

import (
	"log"
	"os"
	"path/filepath"
)

type LocationEnum string

const (
	Database LocationEnum = "database"
)

var relativeLocations = map[LocationEnum]string{
	Database: "data/honcc_0.0.1.db",
}

var locations = map[LocationEnum]string{}

const appRootDirName = "honcc"

func createParentDirOnDisk(path string) {
	locationDir := filepath.Dir(path)

	mkdirErr := os.MkdirAll(locationDir, os.ModePerm)
	if mkdirErr != nil {
		log.Fatalf("Could not create parent dir for location: %s", path)
	}
}

func initLocations() {
	if len(locations) != 0 {
		return
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Could not fetch user home dir...")
		return
	}

	for name, relLocation := range relativeLocations {
		absoluteLocation := filepath.Join(homeDir, appRootDirName, relLocation)
		createParentDirOnDisk(absoluteLocation)
		locations[name] = absoluteLocation
	}
}

func Get(location LocationEnum) string {
	initLocations()
	return locations[location]
}
