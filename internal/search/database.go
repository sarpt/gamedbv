package search

import (
	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/pkg/db"
)

func gamesDetailsFromDatabase(dbConf config.Database, settings Settings, serialNumbers []string) (db.GamesResult, error) {
	var gamesResult db.GamesResult

	database, err := db.OpenDatabase(dbConf)
	defer database.Close()

	if err != nil {
		return gamesResult, err
	}

	gamesQuery := database.NewGamesQuery()
	if len(serialNumbers) > 0 {
		gamesQuery.FilterSerialNumbers(serialNumbers)
	}

	if len(settings.Regions) > 0 {
		gamesQuery.FilterRegions(settings.Regions)
	}

	gamesQuery.Page(settings.Page)
	gamesQuery.Limit(settings.PageLimit)

	gamesResult = gamesQuery.Get()
	return gamesResult, err
}
