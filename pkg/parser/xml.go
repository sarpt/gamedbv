package parser

import (
	"encoding/xml"
	"os"
)

// ParseSourceFile converts database file contents and returns gamedb game Entries from files
func ParseSourceFile(platformConfig Config, provider ModelProvider) error {
	dbFilePath := platformConfig.Filepath()
	dbFile, err := os.Open(dbFilePath)
	if err != nil {
		return err
	}

	err = xml.NewDecoder(dbFile).Decode(provider.Model())
	return err
}
