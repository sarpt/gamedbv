package index

import (
	"fmt"
	"os"
)

// PrepareIndex handles creating index that will be used for searching purposes
func PrepareIndex(creators map[string]Creator, conf PlatformConfig, games []GameSource) error {
	indexPath := conf.IndexFilepath()
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

	if creator, ok := creators[conf.IndexVariant()]; ok {
		err = creator.CreateIndex(conf.DocType(), indexPath, games)
	} else {
		err = fmt.Errorf("Creator not found for the config")
	}
	return err
}
