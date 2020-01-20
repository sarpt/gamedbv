package parser

import (
	"encoding/xml"
	"os"

	"github.com/sarpt/gamedbv/pkg/gametdb"
)

// Config provides information about files to be parsed by xml
type Config interface {
	DatabaseContentFilePath() (string, error)
}

// ParseDatabaseFile converts database file contents and returns gamedb game Entries from files
func ParseDatabaseFile(platformConfig Config) (*gametdb.Datafile, error) {
	var datafile gametdb.Datafile

	dbFilePath, err := platformConfig.DatabaseContentFilePath()
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
