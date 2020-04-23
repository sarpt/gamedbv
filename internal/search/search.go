package search

import (
	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/pkg/db/queries"
)

// FindGames takes platforms, finds indexes which are available to execute query and executes the query on them, returning game results from database
func FindGames(appConf config.App, params Parameters) (queries.GamesResult, error) {
	var gamesResult queries.GamesResult
	var serialNumbers []string

	if params.Text != "" {
		results, err := resultsFromIndex(appConf, params)
		if err != nil {
			return gamesResult, err
		}

		if len(results.Hits) <= 0 {
			return gamesResult, err
		}

		serialNumbers = getSerialNumbersFromGameHits(results.Hits)
	}

	gamesResult, err := gamesDetailsFromDatabase(appConf.Database(), params, serialNumbers)
	return gamesResult, err
}
