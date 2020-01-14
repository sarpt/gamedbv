package search

import (
	"github.com/sarpt/gamedbv/pkg/index"
	"github.com/sarpt/gamedbv/pkg/index/bleve"
	"github.com/sarpt/gamedbv/pkg/platform"
)

// Execute takes platforms, find indexes which are available to execute query and executes the query on them, returning game results
func Execute(platformVariants []platform.Variant, searchParams index.SearchParameters) (string, error) {
	var configs []index.Config
	for _, plat := range platformVariants {
		configs = append(configs, platform.GetConfig(plat))
	}
	searcher := getSearcher(configs)
	return searcher.Search(searchParams)
}

func getSearcher(configs []index.Config) index.Searcher {
	bleveIndex := bleve.NewSearcher()

	for _, conf := range configs {
		if conf.GetIndexType() == "bleve" {
			bleveIndex.Add(conf)
		}
	}

	return bleveIndex
}
