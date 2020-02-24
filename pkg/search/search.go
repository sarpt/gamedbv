package search

import (
	"fmt"

	"github.com/sarpt/gamedbv/pkg/config"
	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/db/models"
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

	gameDetails, err := getGamesDetails(appConf.Database(), res.Hits)
	if err != nil {
		return "", err
	}

	return prepareOutput(gameDetails, res.IgnoredPlatforms), nil
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

func getGamesDetails(dbConf config.Database, hits []index.GameHit) ([]*models.Game, error) {
	var models []*models.Game

	database, err := db.OpenDatabase(dbConf)
	if err != nil {
		return models, err
	}

	var serialNumbers []string
	for _, hit := range hits {
		serialNumbers = append(serialNumbers, hit.ID)
	}

	models = database.GamesForSerials(serialNumbers)
	return models, err
}
