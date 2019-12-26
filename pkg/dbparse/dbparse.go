package dbparse

import (
	"encoding/xml"
	"io/ioutil"

	"github.com/sarpt/gamedbv/pkg/gamedb"
	"github.com/sarpt/gamedbv/pkg/gamedb/gametdb"
	"github.com/sarpt/gamedbv/pkg/platform"
)

// ParseDatabaseFile converts database file contents and returns gamedb game Entries from files
func ParseDatabaseFile(platformDb platform.Variant) (*gametdb.Datafile, error) {
	var datafile gametdb.Datafile

	platformDbInfo := gamedb.GetDbInfo(platformDb)

	dbFilePath, err := platformDbInfo.GetDatabaseContentFilePath()
	if err != nil {
		return &datafile, err
	}

	dbFile, err := ioutil.ReadFile(dbFilePath)
	if err != nil {
		return &datafile, err
	}

	err = xml.Unmarshal(dbFile, &datafile)
	return &datafile, err
}
