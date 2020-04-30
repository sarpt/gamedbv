package parser

import (
	"encoding/xml"
	"os"
)

// Config provides information about files to be parsed by xml
type Config struct {
	Filepath string
}

// ParseSourceFile converts source file contents and returns gamedb game Entries from files
func ParseSourceFile(cfg Config, provider ModelProvider) error {
	sourceFilePath := cfg.Filepath
	sourceFile, err := os.Open(sourceFilePath)
	if err != nil {
		return err
	}

	err = xml.NewDecoder(sourceFile).Decode(provider.Model())
	return err
}
