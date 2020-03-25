package search

import (
	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/pkg/db/models"
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
