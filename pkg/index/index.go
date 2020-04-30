package index

import (
	"fmt"
	"os"
)

// Config provides index information and settings
type Config struct {
	Filepath string
	Variant  string
	Name     string
	DocType  string
}

// PrepareIndex handles creating index that will be used for searching purposes
func PrepareIndex(creators map[string]Creator, cfg Config, games []GameSource) error {
	indexPath := cfg.Filepath
	indexFile, err := os.Stat(indexPath)
	if !os.IsNotExist(err) && err != nil {
		return err
	}

	if indexFile != nil {
		err = os.RemoveAll(indexPath)
		if err != nil {
			return err
		}
	}

	if creator, ok := creators[cfg.Variant]; ok {
		err = creator.CreateIndex(indexPath, games)
	} else {
		err = fmt.Errorf("Creator not found for the config")
	}
	return err
}
