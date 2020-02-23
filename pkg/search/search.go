package search

import (
	"fmt"

	"github.com/sarpt/gamedbv/pkg/config"
	"github.com/sarpt/gamedbv/pkg/index"
	"github.com/sarpt/gamedbv/pkg/index/bleve"
)

// Execute takes platforms, find indexes which are available to execute query and executes the query on them, returning game results
func Execute(appConf config.App, settings Settings) (string, error) {
	searcher := getSearcher(appConf, settings)
	searchParams := mapToSearcherParameters(settings)

	res, err := searcher.Search(searchParams)
	if err != nil {
		return "", err
	}

	return prepareOutput(res), nil
}

func getSearcher(appConf config.App, settings Settings) index.Searcher {
	var configs []index.PlatformConfig
	for _, plat := range settings.Platforms {
		configs = append(configs, appConf.Platform(plat))
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

func prepareOutput(res index.Result) string {
	var out string

	for _, ignored := range res.IgnoredPlatforms {
		out = out + fmt.Sprintf("Search could not be executed for platform %s\n", ignored)
	}

	for _, game := range res.Hits {
		out = out + fmt.Sprintf("[%s] %s\n", game.ID, game.Name)
	}

	return out
}

// func getGameDetails(hit index.GameHit) []models.Game {
// database, err := db.GetDatabase()
// }
