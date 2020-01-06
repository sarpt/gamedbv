package dbindex

import (
	"os"

	"github.com/sarpt/gamedbv/pkg/dbindex/bleve"
	"github.com/sarpt/gamedbv/pkg/dbindex/shared"
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

	err = bleve.CreateIndex(indexPath, games)
	return err
}

// Search takes platforms, find indexes which are available to execute query and executes the query on them, returning game results
func Search(platforms []platform.Variant, searchParams shared.SearchParameters) (string, error) {
	index := getAggregatedIndex(platforms)
	return index.Search(searchParams)
}

func getAggregatedIndex(platforms []platform.Variant) shared.AggregatedIndex {
	bleveIndex := bleve.New()

	for _, plat := range platforms {
		conf := platform.GetConfig(plat)

		if conf.IndexType == "bleve" {
			bleveIndex.Add(conf)
		}
	}

	return bleveIndex
}
