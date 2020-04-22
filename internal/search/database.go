package search

import (
	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/db/queries"
)

func gamesDetailsFromDatabase(dbConf db.Config, settings Settings, serialNumbers []string) (queries.GamesResult, error) {
	var gamesResult queries.GamesResult

	database, err := db.OpenDatabase(dbConf)
	defer database.Close()

	if err != nil {
		return gamesResult, err
	}

	gamesQuery := database.NewGamesQuery()
	if len(serialNumbers) > 0 {
		gamesQuery.FilterUIDs(serialNumbers)
	}

	if len(settings.Regions) > 0 {
		gamesQuery.FilterRegions(settings.Regions)
	}

	gamesQuery.Page(settings.Page)
	gamesQuery.Limit(settings.PageLimit)

	gamesResult = gamesQuery.Get()
	return gamesResult, err
}
