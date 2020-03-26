package search

import (
	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/pkg/index"
	"github.com/sarpt/gamedbv/pkg/index/bleve"
)

func getSearcher(appConf config.App, settings Settings) index.Searcher {
	var configs []index.PlatformConfig
	for _, plat := range settings.Platforms {
		configs = append(configs, appConf.Platform(plat))
	}

	bleveIndex, _ := bleve.NewSearcher(configs)

	return bleveIndex
}

func resultsFromIndex(appConf config.App, settings Settings) (index.Result, error) {
	searcher := getSearcher(appConf, settings)
	defer searcher.Close()

	searchParams := mapToSearcherParameters(settings)

	res, err := searcher.Search(searchParams)
	return res, err
}

func mapToSearcherParameters(settings Settings) index.SearchParameters {
	var platforms []string
	for _, plat := range settings.Platforms {
		platforms = append(platforms, plat.String())
	}

	return index.SearchParameters{
		Text:      settings.Text,
		Platforms: platforms,
	}
}

func getSerialNumbersFromGameHits(hits []index.GameHit) []string {
	var serialNumbers []string
	for _, hit := range hits {
		serialNumbers = append(serialNumbers, hit.ID)
	}

	return serialNumbers
}
