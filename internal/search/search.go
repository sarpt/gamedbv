package search

import (
	"fmt"

	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/db/models"
	"github.com/sarpt/gamedbv/pkg/index"
	"github.com/sarpt/gamedbv/pkg/index/bleve"
)

// Execute takes platforms, finds indexes which are available to execute query and executes the query on them, returning game results from database
func Execute(appConf config.App, settings Settings) (string, error) {
	results, err := resultsFromIndex(appConf, settings)
	if err != nil {
		return "", err
	}

	gameDetails, err := gamesDetailsFromDatabase(appConf.Database(), settings, results.Hits)
	if err != nil {
		return "", err
	}

	return prepareOutput(gameDetails, results.IgnoredPlatforms), nil
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

func prepareOutput(games []*models.Game, ignoredPlatforms []string) string {
	var out string

	for _, ignored := range ignoredPlatforms {
		out = out + fmt.Sprintf("Search could not be executed for platform %s\n", ignored)
	}

	for _, game := range games {
		for _, description := range game.Descriptions {
			if description.Language.Code == "EN" {
				out = out + fmt.Sprintf("===\n[%s] %s\nSynopsis: %s\n", game.SerialNo, description.Title, description.Synopsis)
			}
		}
	}

	return out
}

func gamesDetailsFromDatabase(dbConf config.Database, settings Settings, hits []index.GameHit) ([]*models.Game, error) {
	var models []*models.Game

	database, err := db.OpenDatabase(dbConf)
	if err != nil {
		return models, err
	}

	var serialNumbers []string
	for _, hit := range hits {
		serialNumbers = append(serialNumbers, hit.ID)
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
