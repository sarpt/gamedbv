package dbparse

import (
	"encoding/xml"
	"os"

	"github.com/sarpt/gamedbv/pkg/gametdb"
	"github.com/sarpt/gamedbv/pkg/platform"
)

// ParseDatabaseFile converts database file contents and returns gamedb game Entries from files
func ParseDatabaseFile(variant platform.Variant) (*gametdb.Datafile, error) {
	var datafile gametdb.Datafile

	platformConfig := platform.GetConfig(variant)

	dbFilePath, err := platformConfig.GetDatabaseContentFilePath()
	if err != nil {
		return &datafile, err
	}

	dbFile, err := os.Open(dbFilePath)
	if err != nil {
		return &datafile, err
	}

	err = xml.NewDecoder(dbFile).Decode(&datafile)
	return &datafile, err
}
