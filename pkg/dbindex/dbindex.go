package dbindex

import (
	"os"

	"github.com/sarpt/gamedbv/pkg/gametdb"
	"github.com/sarpt/gamedbv/pkg/platform"
)

// PrepareIndex handles creating index that will be used for searching purposes
func PrepareIndex(variant platform.Variant, games []gametdb.Game) error {
	platformConfig := platform.GetConfig(variant)

	indexPath, err := platformConfig.GetIndexFilePath()
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

	err = createIndex(variant.String(), indexPath, games)
	return err
}
