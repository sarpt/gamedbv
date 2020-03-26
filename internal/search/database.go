package search

import (
	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/db/models"
)

func gamesDetailsFromDatabase(dbConf config.Database, settings Settings, serialNumbers []string) ([]*models.Game, error) {
	var models []*models.Game

	database, err := db.OpenDatabase(dbConf)
	defer database.Close()

	if err != nil {
		return models, err
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

	models, _ = gamesQuery.Get()
	return models, err
}
