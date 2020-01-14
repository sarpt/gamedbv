package index

import (
	"fmt"
	"os"

	"github.com/sarpt/gamedbv/pkg/gametdb"
)

// PrepareIndex handles creating index that will be used for searching purposes
func PrepareIndex(creators map[string]Creator, conf Config, games []gametdb.Game) error {
	indexPath, err := conf.GetIndexFilePath()
	if err != nil {
		return err
	}

	indexFile, err := os.Stat(indexPath)
	if !os.IsNotExist(err) && err != nil {
		return err
	}

	if indexFile != nil {
		err = os.Remove(indexPath)
		if err != nil {
			return err
		}
	}

	if creator, ok := creators[conf.GetIndexType()]; ok {
		err = creator.CreateIndex(indexPath, games)
	} else {
		err = fmt.Errorf("Creator not found for the config")
	}
	return err
}
