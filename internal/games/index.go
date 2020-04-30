package games

import (
	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/pkg/index"
	"github.com/sarpt/gamedbv/pkg/index/bleve"
)

func getSearcher(appConf config.App, params SearchParameters) index.Searcher {
	var configs []index.PlatformConfig
	for _, plat := range params.Platforms {
		configs = append(configs, appConf.Platform(plat))
	}

	bleveIndex, _ := bleve.NewBleveSearcher(configs)

	return bleveIndex
}

func resultsFromIndex(appConf config.App, params SearchParameters) (index.Result, error) {
	searcher := getSearcher(appConf, params)
	defer searcher.Close()

	searchParams := mapToSearcherParameters(params)

	res, err := searcher.Search(searchParams)
	return res, err
}

func mapToSearcherParameters(params SearchParameters) index.SearchParameters {
	var platforms []string
	for _, plat := range params.Platforms {
		platforms = append(platforms, plat.ID())
	}

	return index.SearchParameters{
		Text:      params.Text,
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
