package parser

import (
	"encoding/xml"
	"os"
)

// ParseDatabaseFile converts database file contents and returns gamedb game Entries from files
func ParseDatabaseFile(platformConfig Config, provider ModelProvider) error {
	dbFilePath, err := platformConfig.Filepath()
	if err != nil {
		return err
	}

	dbFile, err := os.Open(dbFilePath)
	if err != nil {
		return err
	}

	err = xml.NewDecoder(dbFile).Decode(provider.Model())
	return err
}
