package games

import (
	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/db/queries"
	"github.com/sarpt/gamedbv/pkg/index"
	"github.com/sarpt/gamedbv/pkg/platform"
)

// Config instructs how to access database and platform indexes
type Config struct {
	Database db.Config
	Indexes  map[platform.Variant]index.Config
}

// Search takes platforms, finds indexes which are available to execute query and executes the query on them, returning game results from database
func Search(conf Config, params SearchParameters) (queries.GamesResult, error) {
	var gamesResult queries.GamesResult
	var serialNumbers []string

	if params.Text != "" {
		results, err := resultsFromIndex(conf, params)
		if err != nil {
			return gamesResult, err
		}

		if len(results.Hits) <= 0 {
			return gamesResult, err
		}

		serialNumbers = getSerialNumbersFromGameHits(results.Hits)
	}

	gamesResult, err := gamesDetailsFromDatabase(conf.Database, params, serialNumbers)
	return gamesResult, err
}
