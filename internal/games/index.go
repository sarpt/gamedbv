package games

import (
	"github.com/sarpt/gamedbv/pkg/index"
	"github.com/sarpt/gamedbv/pkg/index/bleve"
)

func getSearcher(conf Config, params SearchParameters) index.Searcher {
	var configs []index.Config
	for _, plat := range params.Platforms {
		if indexConfig, ok := conf.Indexes[plat]; ok { // not sure how to deal with not ok rn, to be fixed
			configs = append(configs, indexConfig)
		}
	}

	bleveIndex, _ := bleve.NewBleveSearcher(configs)

	return bleveIndex
}

func resultsFromIndex(conf Config, params SearchParameters) (index.Result, error) {
	searcher := getSearcher(conf, params)
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
