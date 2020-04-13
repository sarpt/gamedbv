package search

import (
	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/pkg/db/queries"
)

// FindGames takes platforms, finds indexes which are available to execute query and executes the query on them, returning game results from database
func FindGames(appConf config.App, settings Settings) (queries.GamesResult, error) {
	var gamesResult queries.GamesResult
	var serialNumbers []string

	if settings.Text != "" {
		results, err := resultsFromIndex(appConf, settings)
		if err != nil {
			return gamesResult, err
		}

		if len(results.Hits) <= 0 {
			return gamesResult, err
		}

		serialNumbers = getSerialNumbersFromGameHits(results.Hits)
	}

	gamesResult, err := gamesDetailsFromDatabase(appConf.Database(), settings, serialNumbers)
	return gamesResult, err
}
