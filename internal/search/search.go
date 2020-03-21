package search

import (
	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/db/models"
	"github.com/sarpt/gamedbv/pkg/index"
	"github.com/sarpt/gamedbv/pkg/index/bleve"
)

// FindGames takes platforms, finds indexes which are available to execute query and executes the query on them, returning game results from database
func FindGames(appConf config.App, settings Settings) ([]*models.Game, error) {
	var games []*models.Game
	var serialNumbers []string

	if settings.Text != "" {
		results, err := resultsFromIndex(appConf, settings)
		if err != nil {
			return games, err
		}

		if len(results.Hits) <= 0 {
			return games, err
		}

		serialNumbers = getSerialNumbersFromGameHits(results.Hits)
	}

	games, err := gamesDetailsFromDatabase(appConf.Database(), settings, serialNumbers)
	return games, err
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

func gamesDetailsFromDatabase(dbConf config.Database, settings Settings, serialNumbers []string) ([]*models.Game, error) {
	var models []*models.Game

	database, err := db.OpenDatabase(dbConf)
	if err != nil {
		return models, err
	}

	gamesQuery := database.NewGameQuery()
	if len(serialNumbers) > 0 {
		gamesQuery.FilterSerialNumbers(serialNumbers)
	}

	if len(settings.Regions) > 0 {
		gamesQuery.FilterRegions(settings.Regions)
	}

	models = gamesQuery.Get()
	return models, err
}

func resultsFromIndex(appConf config.App, settings Settings) (index.Result, error) {
	searcher := getSearcher(appConf, settings)
	searchParams := mapToSearcherParameters(settings)

	res, err := searcher.Search(searchParams)
	return res, err
}
