package search

import (
	"github.com/sarpt/gamedbv/pkg/index"
	"github.com/sarpt/gamedbv/pkg/index/bleve"
	"github.com/sarpt/gamedbv/pkg/platform"
)

// Settings allow to control search execution
type Settings struct {
	Platforms []platform.Variant
	Text      string
	Regions   []string
}

// Execute takes platforms, find indexes which are available to execute query and executes the query on them, returning game results
func Execute(settings Settings) (string, error) {
	searcher := getSearcher(settings)
	searchParams := mapToSearcherParameters(settings)

	return searcher.Search(searchParams)
}

func getSearcher(settings Settings) index.Searcher {
	var configs []index.Config
	for _, plat := range settings.Platforms {
		configs = append(configs, platform.GetConfig(plat))
	}

	bleveIndex, _ := bleve.NewSearcher(configs)

	return bleveIndex
}

func mapToSearcherParameters(settings Settings) index.SearchParameters {
	var platforms []string
	for _, plat := range settings.Platforms {
		platforms = append(platforms, plat.String())
	}

	return index.SearchParameters{
		Text:      settings.Text,
		Regions:   settings.Regions,
		Platforms: platforms,
	}
}
